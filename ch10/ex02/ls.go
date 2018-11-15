package main

import (
	"fmt"
	"log"
	"os"

	"./archive"
	_ "./archive/tar"
	_ "./archive/zip"
)

func main() {
	entries, err := archive.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name)
	}
}
