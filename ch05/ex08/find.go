package main

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		return match(n, id)
	}
	return forEachNode(doc, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	willContinue := true
	if pre != nil {
		willContinue = pre(n)
	}

	if !willContinue {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		end := forEachNode(c, pre, post)
		if end != nil {
			return end
		}
	}

	if post != nil {
		post(n)
	}

	return nil
}

func match(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}
