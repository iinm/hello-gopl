package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		in   string
		want []string
	}{
		{
			in: `
			<html>
			<head>
				<link type="text/css" rel="stylesheet" href="style.css">
			</head>
			<body>
				<a href ="1"></a>
				<div>
					<a href ="2"></a>
				</div>
				<a href ="3"></a>
				<img src="image.png">
				<script type="text/javascript" src="script.js"></script>
			</body>
			</html>
			`,
			want: []string{"style.css", "1", "2", "3", "image.png", "script.js"},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(nil, %v)", test.in)
		node, _ := html.Parse(strings.NewReader(test.in))
		got := visit(nil, node)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
