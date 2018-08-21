package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in: `
			<!DOCTYPE html>
			<html>
			<head>
			<link type="text/css" rel="stylesheet" href="style.css">
			<style> body { color: black; } </style>
			<!-- this is comment -->
			</head>
			<body>
			Hello, World!
			<a href ="1">link1</a>
			<pre>こんにちは、世界</pre>
			<div>
			<a href ="2">link2</a>
			</div>
			<a href ="3"></a>
			<img src="image.png">
			<script    type="text/javascript" src="script.js"></script>
			<script type="text/javascript">
			alert("hello")
			</script>
			</body>
			</html>
			`,
			want: `<!DOCTYPE html>
<html>
  <head>
    <link type="text/css" rel="stylesheet" href="style.css"/>
    <style>
    body { color: black; }
    </style>
<!--  this is comment  -->
  </head>
  <body>
  Hello, World!
    <a href="1">
    link1
    </a>
    <pre>こんにちは、世界</pre>
    <div>
      <a href="2">
      link2
      </a>
    </div>
    <a href="3"/>
    <img src="image.png"/>
    <script type="text/javascript" src="script.js"/>
    <script type="text/javascript">
    alert("hello")
    </script>
  </body>
</html>
`,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("forEachNode(%v, ...)", test.in)

		out = new(bytes.Buffer)
		node, _ := html.Parse(strings.NewReader(test.in))

		forEachNode(node, 0, startElement, endElement)
		got := out.(*bytes.Buffer).String()

		// 期待通りにprettifyできていること
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}

		// prettifyしたHTMLをパースできること
		// todo 無効なHTMLを出力しても失敗しない
		node, err := html.Parse(strings.NewReader(got))
		if err != nil {
			t.Errorf("%s = %v, could not parse: %v", descr, got, err)
		}

		if hasErrorNode(node) {
			t.Errorf("%s = %v, should not contain error node", descr, got)
		}
	}
}

func hasErrorNode(n *html.Node) bool {
	if n.Type == html.ErrorNode {
		return true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hasErr := hasErrorNode(c)
		if hasErr {
			return true
		}
	}

	return false
}
