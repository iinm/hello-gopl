package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func read(c *net.TCPConn, out chan<- string) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		out <- scanner.Text()
	}
	c.CloseRead()
	close(out)
}

func handleConn(c *net.TCPConn) {
	var echoWaitGroup sync.WaitGroup
	input := make(chan string)

	go read(c, input)

loop:
	for {
		timer := time.NewTimer(10 * time.Second)
		select {
		case <-timer.C:
			break loop
		case s, ok := <-input:
			if !ok {
				break loop
			}
			echoWaitGroup.Add(1)
			go echo(c, s, 1*time.Second, &echoWaitGroup)
		}
		timer.Stop()
	}

	go func() {
		echoWaitGroup.Wait()
		c.CloseWrite()
	}()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		tcpConn := conn.(*net.TCPConn)
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(tcpConn)
	}
}
