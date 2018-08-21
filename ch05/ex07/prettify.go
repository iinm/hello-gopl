package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer

func main() {
	out = os.Stdout
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "prettify: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, 0, startElement, endElement)
}

func forEachNode(n *html.Node, depth int, pre, post func(n *html.Node, depth int)) {
	nextDepth := depth + 1
	if n.Type == html.DoctypeNode || n.Type == html.DocumentNode {
		depth, nextDepth = 0, 0
	}

	if pre != nil {
		pre(n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, nextDepth, pre, post)
	}

	if post != nil {
		post(n, depth)
	}
}

var ignoreIndent = map[string]bool{
	"pre":      true,
	"textarea": true,
}

func startElement(n *html.Node, depth int) {
	if n.Type == html.DoctypeNode {
		fmt.Fprintf(out, "<!DOCTYPE %s>\n", n.Data)
	}

	if n.Type == html.ElementNode {
		// 小要素がないなら終了タグを省略する
		optSlash := ""
		if n.FirstChild == nil {
			optSlash = "/"
		}
		// 属性を取得
		attrStr := makeAttrStr(n.Attr)
		if len(attrStr) > 0 {
			attrStr = " " + attrStr
		}
		// 改行するか?
		optNewline := "\n"
		if _, ok := ignoreIndent[n.Data]; ok {
			optNewline = ""
		}
		fmt.Fprintf(out, "%*s<%s%s%s>%s", depth*2, "", n.Data, attrStr, optSlash, optNewline)

		// 小要素がテキストなら表示
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			indent := ""
			text := n.FirstChild.Data
			optNewline = ""
			if _, ok := ignoreIndent[n.Data]; !ok {
				text = strings.TrimSpace(n.FirstChild.Data)
				indent = fmt.Sprintf("%*s", depth*2, "")
				optNewline = "\n"
			}
			if len(text) > 0 {
				fmt.Fprintf(out, indent+text+optNewline)
			}
		}
	}

	if n.Type == html.CommentNode {
		//fmt.Printf("%*s<!-- %s -->\n", depth*2, "", strings.TrimSpace(n.Data))
		fmt.Fprintf(out, "<!-- %s -->\n", n.Data)
	}
}

func endElement(n *html.Node, depth int) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		indent := fmt.Sprintf("%*s", depth*2, "")
		if _, ok := ignoreIndent[n.Data]; ok {
			indent = ""
		}
		fmt.Fprintf(out, "%s</%s>\n", indent, n.Data)
	}
}

func makeAttrStr(attr []html.Attribute) string {
	var buf bytes.Buffer
	for i, a := range attr {
		if i > 0 {
			buf.WriteByte(' ')
		}
		escapedVal := strings.Replace(a.Val, `"`, `\"`, -1)
		s := fmt.Sprintf(`%s="%s"`, a.Key, escapedVal)
		buf.WriteString(s)
	}
	return buf.String()
}
