package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	echo(os.Args)
}

func echo(args []string) {
	for i, arg := range args {
		fmt.Fprintf(out, "%d\t%s\n", i, arg)
	}
}
