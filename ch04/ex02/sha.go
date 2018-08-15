package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.Int("mode", 256, "256 or 384 or 512")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	var buf bytes.Buffer
	for input.Scan() {
		buf.Write(input.Bytes())
	}

	switch *mode {
	case 256:
		sum256 := sha256.Sum256(buf.Bytes())
		fmt.Printf("%x\n", sum256)
	case 384:
		sum384 := sha512.Sum384(buf.Bytes())
		fmt.Printf("%x\n", sum384)
	case 512:
		sum512 := sha512.Sum512(buf.Bytes())
		fmt.Printf("%x\n", sum512)
	default:
		fmt.Fprintf(os.Stderr, "invalid mode: %d\n", *mode)
		os.Exit(1)
	}
}
