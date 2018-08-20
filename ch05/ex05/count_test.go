package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordAndImages(t *testing.T) {
	tests := []struct {
		in                        string
		wantedWords, wantedImages int
	}{
		{
			in: `<html>
				 <head><style>body { color: black  }</style></head>
				 <body>
					<a href="1">one</a>
					<div>
						<a href ="2">two</a>
						<p>three</p>
						<img src="foo.png">
						<img src="bar.png">
					</div>
					<a href ="3"></a>
					<script>alert("script!")</script>
				 </body>
				 </html>`,
			wantedWords:  3,
			wantedImages: 2,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countWordAndImages(%v)", test.in)
		node, _ := html.Parse(strings.NewReader(test.in))
		words, images := countWordAndImages(node)
		if words != test.wantedWords || images != test.wantedImages {
			t.Errorf("%s = (%d, %d), want (%d, %d)",
				descr, words, images, test.wantedWords, test.wantedImages)
		}
	}
}
