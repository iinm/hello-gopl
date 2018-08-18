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
			<head></head>
			<body>
				<a href ="1"></a>
				<div>
					<a href ="2"></a>
				</div>
				<a href ="3"></a>
			</body>
			</html>
			`,
			want: []string{"1", "2", "3"},
		},
		{
			in: `
			<html>
			<head></head>
			<body>
				<a href ="1"></a>
				<div>
					<a href ="2"></a>
					<a href ="3"></a>
					<div>
						<a href ="4"></a>
						<a href ="5"></a>
					</div>
					<a href ="6"></a>
					<a href ="7"></a>
				</div>
				<a href ="8"></a>
			</body>
			</html>
			`,
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
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
