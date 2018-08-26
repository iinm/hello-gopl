package links

import (
	"fmt"
	"io"
	"net/url"

	"golang.org/x/net/html"
)

// Extract htmlBodyをパースして、HTMLに含まれるリンクを返す。htmlの取得元のURL uを使って、絶対パスを返す。
func Extract(htmlBody io.Reader, u *url.URL) ([]string, error) {
	doc, err := html.Parse(htmlBody)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %v", err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := u.Parse(a.Val)
				if err != nil {
					continue // 不正なURLを無視
				}
				links = append(links, link.String())
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
