package selector

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	in := `
	<html>
	<head></head>
	<img src="foo.png"/>
	<h1></h1>
	<div>
		<h2></h2>
		<h3></h3>
	</div>
	<div>
		<h2></h2>
		<h3></h3>
		<h3></h3>
	</div>
	</html>
	`
	doc, _ := html.Parse(strings.NewReader(in))

	tests := []struct {
		tags           []string
		wantedTagNames []string
	}{
		{[]string{}, []string{}},
		{[]string{"img"}, []string{"img"}},
		{[]string{"h1", "h3"}, []string{"h1", "h3", "h3", "h3"}},
		{[]string{"h1", "h2", "h3"}, []string{"h1", "h2", "h3", "h2", "h3", "h3"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("ElementByTagName(%v, %q)", in, test.tags)
		got := ElementByTagName(doc, test.tags...)
		if !reflect.DeepEqual(tagNames(got), test.wantedTagNames) {
			t.Errorf("%s = %q, want %q", descr, tagNames(got), test.wantedTagNames)
		}
	}
}

func tagNames(nodes []*html.Node) []string {
	names := []string{}
	for _, n := range nodes {
		names = append(names, n.Data)
	}
	return names
}
