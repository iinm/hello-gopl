package main

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	var targetNode *html.Node
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					targetNode = n
					return false
				}
			}
		}
		return true
	}

	forEachNode(doc, pre, nil)

	return targetNode
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if n == nil {
		return
	}

	if pre != nil {
		willContinue := pre(n)
		if !willContinue {
			return
		}
	}

	forEachNode(n.FirstChild, pre, post)
	// todo: forEachNodeが継続するか否かを返さないとn.NextSiblingもチェックしてしまう
	forEachNode(n.NextSibling, pre, post)

	if post != nil {
		post(n)
	}
}
