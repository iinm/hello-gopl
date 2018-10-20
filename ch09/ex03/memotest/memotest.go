package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://golang.org",
			"https://godoc.org",
			"https://godoc.org",
			"https://play.golang.org",
			"https://play.golang.org",
			"http://gopl.io",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, cancel <-chan struct{}) (interface{}, error)
}

func ConcurrentCancelAll(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			close(cancel)
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

// 奇数番目のリクエストをキャンセルする
func ConcurrentCancelOdd(t *testing.T, m M) {
	var n sync.WaitGroup
	urlCount := 0
	for url := range incomingURLs() {
		urlCount++
		n.Add(1)
		go func(url string, urlCount int) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			if urlCount%2 == 1 {
				close(cancel)
			}
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url, urlCount)
	}
	n.Wait()
}

// 偶数番目のリクエストをキャンセルする
func ConcurrentCancelEven(t *testing.T, m M) {
	var n sync.WaitGroup
	urlCount := 0
	for url := range incomingURLs() {
		urlCount++
		n.Add(1)
		go func(url string, urlCount int) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			if urlCount%2 == 0 {
				close(cancel)
			}
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url, urlCount)
	}
	n.Wait()
}

// waitMillisecだけ待ってからリクエストをキャンセルする
func ConcurrentCancelAllWait(t *testing.T, m M, waitMilliSec time.Duration) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			cancel := make(chan struct{})
			go func() {
				<-time.After(waitMilliSec * time.Millisecond)
				close(cancel)
			}()
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, make(chan struct{}))
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, make(chan struct{}))
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}
