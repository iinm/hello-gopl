package selector

import "golang.org/x/net/html"

func ElementByTagName(n *html.Node, names ...string) []*html.Node {
	nodes := []*html.Node{}
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, n)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		children := ElementByTagName(c, names...)
		nodes = append(nodes, children...)
	}

	return nodes
}
