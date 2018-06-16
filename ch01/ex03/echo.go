package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func slowEcho(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func echo(args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}
