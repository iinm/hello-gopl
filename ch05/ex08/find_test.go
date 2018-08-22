package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	in := `
	<html>
	<head><style>body { color: black  }</style></head>
	<body>
		<a id="one" href="1">one</a>
		<div>
			<a id="two" href ="2">two</a>
			<p id="three">three</p>
			<img src="foo.png">
			<img src="bar.png">
		</div>
		<a href ="3"></a>
		<script>alert("script!")</script>
	</body>
	</html>
	`
	node, _ := html.Parse(strings.NewReader(in))

	tests := []struct {
		id   string
		want *html.Node
	}{
		{id: "zero", want: nil},
		{id: "one", want: &html.Node{Data: "a"}},
		{id: "two", want: &html.Node{Data: "a"}},
		{id: "three", want: &html.Node{Data: "p"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("ElementByID(%v, %s)", in, test.id)
		got := ElementByID(node, test.id)

		if !equal(got, test.want) {
			t.Errorf("%s = %#v, want %#v", descr, got, test.want)
		}
	}
}

func equal(n, m *html.Node) bool {
	if n == nil && m == nil {
		return true
	}

	if n != nil && m != nil && n.Data == n.Data {
		return true
	}

	return false
}
