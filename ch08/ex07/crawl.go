package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	saveDir := flag.String("out", ".", "クロールしたページを保存するディレクトリ")
	nWorkers := flag.Int("workers", 1, "並列ダウンロード数")
	flag.Parse()

	root, err := url.Parse(flag.Args()[0])
	if err != nil {
		panic(err)
	}

	// 同じドメインのページだけ取得する
	linkFilter := func(link *url.URL) bool { return link.Host == root.Host }

	links := make(chan *url.URL) // クロール対象のURL
	go func() { links <- root }()
	unseenLinks := iterUnseenLinks(links)
	responses := crawl(unseenLinks, links, linkFilter, *nWorkers)
	saveDirAbs, err := filepath.Abs(*saveDir)
	if err != nil {
		log.Fatal(err)
	}
	save(responses, saveDirAbs)
}

func iterUnseenLinks(links <-chan *url.URL) <-chan *url.URL {
	unseenLinks := make(chan *url.URL)
	go func() {
		seen := make(map[string]bool)
		for link := range links {
			if !seen[link.String()] {
				seen[link.String()] = true
				unseenLinks <- link
			}
		}
		close(unseenLinks)
	}()
	return unseenLinks
}

type response struct {
	link *url.URL
	body []byte
	err  error
}

func crawl(links <-chan *url.URL, nextLinks chan<- *url.URL, linkFilter func(*url.URL) bool, worker int) <-chan *response {
	responses := make(chan *response)
	var wg sync.WaitGroup
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range links {
				resp := request(link)
				if !linkFilter(link) {
					// Note: リダイレクトで別ドメインのリソースをGETする場合がある
					// レスポンスを破棄して、リンクも辿らない
					break
				}
				go func(resp *response) {
					foundLinks, err := extractLinks(resp.link, resp.body)
					if err != nil {
						log.Printf("error: %v\n", err)
						return
					}
					for _, l := range foundLinks {
						if linkFilter(l) {
							nextLinks <- l
						}
					}
				}(resp)
				responses <- resp
			}
		}()
		go func() {
			wg.Wait()
			close(responses)
		}()
	}
	return responses
}

func save(responses <-chan *response, dir string) {
	for resp := range responses {
		if resp.err != nil {
			log.Printf("warning: skip save: %v\n", resp.err)
			continue
		}
		filename := url2filepath(resp.link, dir)
		log.Printf("info: save %q --> %q\n", resp.link.String(), filename)
		dirname := filepath.Dir(filename)
		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			if err := os.MkdirAll(dirname, 0744); err != nil {
				log.Printf("error: %v\n", err)
				continue
			}
		}
		modifiedBody, err := modifyLinks(resp.link, resp.body, dir)
		if err != nil {
			log.Printf("error: %v\n", err)
			continue
		}
		if err := ioutil.WriteFile(filename, modifiedBody, 0644); err != nil {
			log.Printf("error: %v\n", err)
		}
	}
}

func url2filepath(link *url.URL, dir string) string {
	filename := filepath.Join(dir, link.Host, link.Path)
	if !strings.HasSuffix(link.Path, ".html") {
		filename = filepath.Join(filename, "index.html")
	}
	return filename
}

func request(link *url.URL) *response {
	resp, err := http.Get(link.String())
	if err != nil {
		return &response{link: link, err: err}
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &response{link: link, err: fmt.Errorf("getting %s: %s", link, resp.Status)}
	}
	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return &response{link: link, err: err}
	}
	return &response{link: resp.Request.URL, body: buf.Bytes()}
}

func modifyLinks(link *url.URL, body []byte, dir string) ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %v", err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			modfiedAttributes := []html.Attribute{}
			for _, a := range n.Attr {
				if a.Key == "href" {
					l, err := link.Parse(a.Val)
					if err != nil {
						log.Printf("warning: %v\n", err)
						continue // ignore bad URLs
					}
					a.Val = url2filepath(l, dir)
					//fmt.Printf("%s  %s\n", link, a.Val)
				}
				modfiedAttributes = append(modfiedAttributes, a)
			}
			n.Attr = modfiedAttributes
		}
	}
	forEachNode(doc, visitNode, nil)

	buf := &bytes.Buffer{}
	if err := html.Render(buf, doc); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func extractLinks(link *url.URL, body []byte) ([]*url.URL, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %v", err)
	}
	var links []*url.URL
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				l, err := link.Parse(a.Val)
				if err != nil {
					log.Printf("error: %v\n", err)
					continue // ignore bad URLs
				}
				links = append(links, l)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
