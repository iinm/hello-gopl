package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 4)
	cancel := make(chan struct{})
	defer close(cancel)
	go func() { responses <- request("http://ftp.tw.debian.org/debian/README", cancel) }()
	go func() { responses <- request("http://ftp.is.debian.org/debian/README", cancel) }()
	go func() { responses <- request("http://ftp.ua.debian.org/debian/README", cancel) }()
	go func() { responses <- request("http://ftp.cl.debian.org/debian/README", cancel) }()
	return <-responses
}

func request(url string, cancel <-chan struct{}) (response string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	req.Cancel = cancel
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer res.Body.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(res.Body)
	log.Printf("got response from: %s", url)
	return buf.String()
}
