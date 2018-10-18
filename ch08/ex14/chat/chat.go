package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	ch   chan<- string // 送信用メッセージチャネル
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントからのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) // 接続中のクライアント
	for {
		select {
		case msg := <-messages:
			// 受信したメッセージを全てのクライアントの送信用チャネルへブロードキャストする
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			//.現在のクライアントの集まりを通知
			cli.ch <- fmt.Sprintf("Active clients: %s", stringifyClients(clients))
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func stringifyClients(clients map[client]bool) string {
	names := make([]string, 0)
	for c, _ := range clients {
		names = append(names, c.name)
	}
	return strings.Join(names, ", ")
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	reader := bufio.NewReader(conn)
	var name string
	for len(name) == 0 {
		ch <- "Input your name:"
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		name = string(line)
	}

	clientAddress := conn.RemoteAddr().String()
	c := client{ch, name}
	ch <- fmt.Sprintf("You are %s (%s)", c.name, clientAddress)
	messages <- c.name + " has arrived"
	entering <- c

	clientInput := make(chan string)
	go func() {
		input := bufio.NewScanner(reader)
		for input.Scan() {
			clientInput <- input.Text()
		}
		close(clientInput)
		// NOTE: ignoring potential errors from input.Err()
	}()

loop:
	for {
		//timer := time.NewTimer(10 * time.Second)
		timer := time.NewTimer(5 * time.Minute)
		select {
		case input, ok := <-clientInput:
			if !ok {
				break loop
			}
			messages <- c.name + ": " + input
		case <-timer.C:
			break loop
		}
		timer.Stop()
	}

	leaving <- c
	messages <- c.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
