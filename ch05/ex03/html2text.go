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

	texts := getTexts(nil, doc)

	for _, text := range texts {
		fmt.Printf("%s\n", text)
	}
}

func getTexts(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}
		texts = getTexts(texts, c)
	}
	return texts
}
