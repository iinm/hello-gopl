package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

func main() {
	times := make(chan time)

	for _, arg := range os.Args[1:] {
		labelURLs := strings.Split(arg, "=") // e.g. Tokyo=localhost:8020
		label, url := labelURLs[0], labelURLs[1]
		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatal(err)
		}
		go readTime(conn, label, times)
	}

	label2time := make(map[string]string)
	for t := range times {
		label2time[t.label] = t.time
		fmt.Printf("\r%s", formatTimes(label2time))
	}
}

func readTime(conn net.Conn, label string, times chan<- time) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		times <- time{input.Text(), label}
	}
}

func formatTimes(label2time map[string]string) string {
	labels := []string{}
	for l, _ := range label2time {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	buf := &bytes.Buffer{}
	for _, l := range labels {
		fmt.Fprintf(buf, "| %s: %s |", l, label2time[l])
	}
	return buf.String()
}

type time struct {
	time  string
	label string
}
