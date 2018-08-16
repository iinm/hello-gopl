package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Item struct {
	Num        int
	SafeTitle  string `json:"safe_title"`
	Transcript string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: go run xkcd.go ./data word\n")
		os.Exit(1)
	}

	dataDir := os.Args[1]
	word := strings.ToLower(os.Args[2])

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		raw, err := ioutil.ReadFile(path.Join(dataDir, f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		var item Item
		json.Unmarshal(raw, &item)

		if strings.Index(strings.ToLower(item.SafeTitle), word) > -1 || strings.Index(strings.ToLower(item.Transcript), word) > -1 {
			fmt.Printf("\n--- https://xkcd.com/%d/\n", item.Num)
			fmt.Printf("title: %s\n", item.SafeTitle)
			fmt.Printf("transcript: %s\n", item.Transcript)
		}
	}
}
