package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type linkDepth struct {
	link  string
	depth int
}

var tokens = make(chan struct{}, 20)

func crawl(link linkDepth, maxDepth int) []linkDepth {
	fmt.Printf("depth: %d\turl: %s\n", link.depth, link.link)
	if link.depth == maxDepth {
		return nil
	}

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(link.link)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	return bundleDepth(list, link.depth+1)
}

func main() {
	maxDepth := flag.Int("depth", 1, "depth")
	flag.Parse()

	worklist := make(chan []linkDepth)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- bundleDepth(flag.Args(), 0) }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.link] {
				seen[link.link] = true
				n++
				go func(link linkDepth) {
					worklist <- crawl(link, *maxDepth)
				}(link)
			}
		}
	}
}

func bundleDepth(links []string, depth int) []linkDepth {
	linksWithDepth := make([]linkDepth, len(links))
	for _, l := range links {
		linksWithDepth = append(linksWithDepth, linkDepth{l, depth})
	}
	return linksWithDepth
}
