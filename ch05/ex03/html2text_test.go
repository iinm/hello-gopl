package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountTags(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{
			in: `<html>` +
				`<head><style>body { color: black  }</style></head>` +
				`<body>` +
				`<a href="1">one</a><div><a href ="2">two</a><p>three</p></div><a href ="3"></a>` +
				`<script>alert("script!")</script>` +
				`</body>` +
				`</html>`,
			want: []string{"one", "two", "three"},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("getTexts(nil, %v)", test.in)
		node, _ := html.Parse(strings.NewReader(test.in))
		texts := getTexts(nil, node)
		if !reflect.DeepEqual(texts, test.want) {
			t.Errorf("%s = %q, want %q", descr, texts, test.want)
		}
	}
}
