package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string) // line -> locations (filename:linuNumber)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "STDIN", counts, locations)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, arg, counts, locations)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, locations[line])
		}
	}
}

func countLines(f io.Reader, filename string, counts map[string]int, locations map[string][]string) {
	input := bufio.NewScanner(f)
	lineNumber := 0
	for input.Scan() {
		lineNumber++
		line := input.Text()
		counts[line]++
		location := fmt.Sprintf("%s:%d", filename, lineNumber)
		locations[line] = append(locations[line], location)
	}
	// 注意: input.Err() からのエラーの可能性を無視している
}
