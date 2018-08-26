package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"./links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func urlToFilepath(u *url.URL) (dirname, filename string) {
	dir, file := path.Split(u.Path)
	ext := filepath.Ext(file)
	if ext == "" {
		dir = u.Path
		file = "index.html"
	}
	query, flagment := "", ""
	if len(u.RawQuery) > 0 {
		query = "?" + u.RawQuery
	}
	if len(u.Fragment) > 0 {
		flagment = "#" + u.Fragment
	}

	dirname = filepath.Join(u.Host, dir)
	filename = file + query + flagment
	return dirname, filename
}

func save(body io.Reader, u *url.URL) {
	dirname, filename := urlToFilepath(u)
	if _, err := os.Stat(dirname); err != nil {
		err := os.MkdirAll(dirname, os.ModePerm)
		if err != nil {
			log.Print(err)
			return
		}
	}

	f, err := os.Create(filepath.Join(dirname, filename))
	if err != nil {
		log.Print(err)
		return
	}

	_, err = io.Copy(f, body)
	if err != nil {
		log.Print(err)
	}
}

func crawl(rawurl string) []string {
	fmt.Println(rawurl)

	resp, err := http.Get(rawurl)
	if err != nil {
		log.Print(err)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Printf("getting %s: %s", rawurl, resp.Status)
		return nil
	}

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)
	save(tee, resp.Request.URL)

	list, err := links.Extract(&buf, resp.Request.URL)
	resp.Body.Close()
	if err != nil {
		log.Print(err)
	}

	// 同じドメインに絞る
	var filteredURLs []string
	for _, raw := range list {
		u, err := url.Parse(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		if u.Hostname() != resp.Request.URL.Hostname() {
			continue
		}
		filteredURLs = append(filteredURLs, raw)
	}

	return filteredURLs
}
