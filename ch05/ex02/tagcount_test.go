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
		want map[string]int
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
			want: map[string]int{
				"html": 1,
				"head": 1,
				"body": 1,
				"a":    3,
				"div":  1,
			},
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
			want: map[string]int{
				"html": 1,
				"head": 1,
				"body": 1,
				"a":    8,
				"div":  2,
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countTags([], %v)", test.in)
		node, _ := html.Parse(strings.NewReader(test.in))
		counts := make(map[string]int)
		countTags(counts, node)
		if !reflect.DeepEqual(counts, test.want) {
			t.Errorf("%s = %v, want %v", descr, counts, test.want)
		}
	}
}
