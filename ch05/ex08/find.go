package main

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	var targetNode *html.Node
	pre := func(n *html.Node) bool {
		willContinue := true
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					targetNode = n
					willContinue = false
				}
			}
		}
		return willContinue
	}

	forEachNode(doc, pre, nil)

	return targetNode
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	willContinue := true
	if pre != nil {
		willContinue = pre(n)
	}

	if !willContinue {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
