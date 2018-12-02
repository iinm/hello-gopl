package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	reader := bufio.NewReader(in)
	for {
		r, n, err := reader.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(out, "rune\tcount\n")
	for c, n := range counts {
		fmt.Fprintf(out, "%q\t%d\n", c, n)
	}
	fmt.Fprint(out, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(out, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(out, "\n%d invalid UTF-8 characters\n", invalid)
	}
}
