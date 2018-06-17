package main

import (
	"fmt"
	"io"
	"strings"
)

func slowEcho(args []string, out io.Writer) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func echo(args []string, out io.Writer) {
	fmt.Fprintln(out, strings.Join(args, " "))
}
