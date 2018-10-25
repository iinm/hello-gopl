package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

type client struct {
	conn        net.Conn
	dataConn    net.Conn
	dataAddress string
	username    string
	authorized  bool
}

func (c *client) request(cmd string, args []string) {
	log.Printf("cmd: %q, args: %v", cmd, args)

	switch strings.ToLower(cmd) {
	case "user":
		if len(args) != 1 {
			reply(c.conn, "501 Syntax error in parameters or arguments")
			return
		}
		c.username = args[0]
		reply(c.conn, "331 Anonymous login ok, send complete email address as your password")

	case "pass":
		c.authorized = true
		reply(c.conn, "230 Anonymous User logged in")

	case "syst":
		reply(c.conn, "215 UNIX")

	case "port":
		if len(args) != 1 {
			reply(c.conn, "501 Syntax error in parameters or arguments")
			return
		}
		numbers := make([]int, 0)
		for _, s := range strings.Split(args[0], ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Print(err)
				reply(c.conn, "501 Syntax error in parameters or arguments")
				return
			}
			numbers = append(numbers, n)
		}
		c.dataAddress = fmt.Sprintf("%d.%d.%d.%d:%d",
			numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]*256+numbers[5])
		reply(c.conn, "200 Command OK")

	case "list":
		reply(c.conn, "150 Opening ASCII mode data connection")
		var err error
		c.dataConn, err = net.Dial("tcp", c.dataAddress)
		if err != nil {
			log.Print(err)
			reply(c.conn, "425 Can't open data connection")
			return
		}

		buf := new(bytes.Buffer)
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			buf.WriteString(f.Name() + "\r\n")
		}

		_, err = c.dataConn.Write(buf.Bytes())
		if err != nil {
			log.Print(err)
		}
		c.dataConn.Close()
		reply(c.conn, "226 Closing data connection")

	case "retr":
		if len(args) != 1 {
			reply(c.conn, "501 Syntax error in parameters or arguments")
			return
		}
		reply(c.conn, "150 Opening ASCII mode data connection")
		var err error
		c.dataConn, err = net.Dial("tcp", c.dataAddress)
		if err != nil {
			log.Print(err)
			reply(c.conn, "425 Can't open data connection")
			return
		}

		content, err := ioutil.ReadFile("./" + args[0])
		if err != nil {
			log.Fatal(err)
		}
		_, err = c.dataConn.Write(content)
		if err != nil {
			log.Print(err)
		}
		c.dataConn.Close()
		reply(c.conn, "226 Closing data connection")

	case "quit":
		reply(c.conn, "221 Goodbye")
		c.conn.Close()

	default:
		reply(c.conn, "502 Command not implemented")
	}
}

func handleConn(conn net.Conn) {
	in := make(chan string)
	cl := &client{conn: conn}
	go recieve(conn, in)
	reply(conn, "220 Service Ready")
	for s := range in {
		ss := strings.Split(s, " ")
		if len(ss) == 0 {
			continue
		}
		cmd := ss[0]
		var args []string
		if len(ss) >= 2 {
			args = ss[1:]
		}
		cl.request(cmd, args)
	}
}

func recieve(conn net.Conn, out chan<- string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%s: %q", conn.RemoteAddr(), line)
		out <- line
	}
	//conn.Close()
}

func reply(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s + "\r\n"))
	if err != nil {
		log.Print(err)
	}
}
