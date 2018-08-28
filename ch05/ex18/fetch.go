package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch(os.Args[1]))
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func(err *error) {
		if closeErr := f.Close(); *err == nil {
			*err = closeErr
		}
	}(&err)

	return local, n, err
}
