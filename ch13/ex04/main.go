package main

import (
	"io"
	"log"
	"os"

	"./bzip"
)

func main() {
	w, err := bzip.NewWriter(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
