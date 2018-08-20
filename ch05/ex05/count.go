package main

import (
	"strings"

	"golang.org/x/net/html"
)

func countWordAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}
		w, i := countWordAndImages(c)
		words += w
		images += i
	}
	return words, images
}
