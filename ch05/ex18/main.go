package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch(os.Args[1]))
}

func fetch(url string) (filename string, n int64, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	filename = path.Base(resp.Request.URL.Path)
	if filename == "/" || filename == "." {
		filename = "index.html"
	}
	var f *os.File
	f, err = os.Create(filename)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return filename, n, err
}
