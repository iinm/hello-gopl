// html内のリンクを表示する
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		var linkAttrKey string
		switch n.Data {
		case "a", "link":
			linkAttrKey = "href"
		case "img", "script":
			linkAttrKey = "src"
		}
		if len(linkAttrKey) > 0 {
			for _, a := range n.Attr {
				if a.Key == linkAttrKey {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
