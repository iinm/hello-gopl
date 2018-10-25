package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
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
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	c, err := NewClient(conn)
	if err != nil {
		log.Print(err)
		fmt.Fprint(conn, "421 Service not available, closing control connection\r\n")
		conn.Close()
		return
	}

	c.reply("220 Service Ready")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%s: %q", conn.RemoteAddr(), line)
		ss := strings.Split(line, " ")
		if len(ss) == 0 {
			c.reply("501 Syntax error in parameters or arguments")
			continue
		}
		cmd := ss[0]
		var args []string
		if len(ss) >= 2 {
			args = ss[1:]
		}
		c.request(cmd, args)
	}
}

type client struct {
	conn        net.Conn
	dataConn    net.Conn
	dataAddress string
	username    string
	authorized  bool
	workdir     string
	initWorkdir string // このディレクトリ配下のみアクセス可能にする
	dataType    string
}

func NewClient(conn net.Conn) (*client, error) {
	workdir, err := filepath.Abs("./")
	if err != nil {
		return nil, err
	}
	c := &client{
		conn:        conn,
		dataType:    "A",
		workdir:     workdir,
		initWorkdir: workdir,
	}
	return c, nil
}

func (c *client) reply(s string) {
	_, err := c.conn.Write([]byte(s + "\r\n"))
	if err != nil {
		log.Print(err)
	}
}

func (c *client) replyf(s string, a ...interface{}) {
	c.reply(fmt.Sprintf(s, a...))
}

func (c *client) transfer(r io.Reader) {
	c.replyf("150 Opening (%s) type data connection", c.dataType)
	var err error
	c.dataConn, err = net.Dial("tcp", c.dataAddress)
	if err != nil {
		log.Print(err)
		c.reply("425 Can't open data connection")
		return
	}
	defer c.dataConn.Close()
	if c.dataType == "A" {
		r = convertToAscii(r)
	}
	_, err = io.Copy(c.dataConn, r)
	if err != nil {
		log.Print(err)
	}
	c.reply("226 Closing data connection")
}

func (c *client) request(cmd string, args []string) {
	log.Printf("cmd: %q, args: %v", cmd, args)

	switch strings.ToLower(cmd) {
	case "user":
		if len(args) != 1 {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		c.username = args[0]
		c.reply("331 Anonymous login ok, send complete email address as your password")

	case "pass":
		c.authorized = true
		c.reply("230 Anonymous User logged in")

	case "syst":
		c.reply("215 UNIX")

	case "port":
		if len(args) != 1 {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		ss := strings.Split(args[0], ",")
		if len(ss) != 6 {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		numbers := make([]int, 0)
		for _, s := range ss {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Print(err)
				c.reply("501 Syntax error in parameters or arguments")
				return
			}
			numbers = append(numbers, n)
		}
		c.dataAddress = fmt.Sprintf("%d.%d.%d.%d:%d",
			numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]*256+numbers[5])
		c.reply("200 Command OK")

	case "pwd":
		c.replyf("257 %q", c.workdir)

	case "cwd":
		if len(args) != 1 {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		dir := args[0]
		if !strings.HasPrefix(dir, "/") {
			dir = filepath.Join(c.workdir, dir)
		}
		abspath, err := filepath.Abs(dir)
		if err != nil {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		if !strings.HasPrefix(abspath, c.initWorkdir) {
			c.reply("550 Requested action not taken. File unavailable")
			return
		}
		if _, err := os.Stat(abspath); os.IsNotExist(err) {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}

		c.workdir = abspath
		c.replyf("200 directory changed to %s", c.workdir)

	case "list":
		dir := c.workdir
		if len(args) == 1 {
			dir = args[0]
		}
		if !strings.HasPrefix(dir, "/") {
			dir = filepath.Join(c.workdir, dir)
		}
		abspath, err := filepath.Abs(dir)
		if err != nil {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		if !strings.HasPrefix(abspath, c.initWorkdir) {
			c.reply("550 Requested action not taken. File unavailable")
			return
		}
		if _, err := os.Stat(abspath); os.IsNotExist(err) {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		buf := new(bytes.Buffer)
		files, err := ioutil.ReadDir(abspath)
		if err != nil {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		for _, f := range files {
			buf.WriteString(f.Name() + "\r\n")
		}
		c.transfer(buf)

	case "retr":
		if len(args) != 1 {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		filename := filepath.Join(c.workdir, args[0])
		abspath, err := filepath.Abs(filename)
		if err != nil {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		if !strings.HasPrefix(abspath, c.initWorkdir) {
			c.reply("550 Requested action not taken. File unavailable")
			return
		}
		if _, err := os.Stat(abspath); os.IsNotExist(err) {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}

		file, err := os.Open(abspath)
		if err != nil {
			log.Print(err)
			c.replyf("550 %s", err)
			return
		}
		defer file.Close()
		c.transfer(file)

	case "type":
		if len(args) != 1 || (args[0] != "A" && args[0] != "I") {
			c.reply("501 Syntax error in parameters or arguments")
			return
		}
		c.dataType = args[0]
		c.replyf("200 data type changed to %s", c.dataType)

	case "quit":
		c.reply("221 Goodbye")
		c.conn.Close()

	default:
		c.reply("502 Command not implemented")
	}
}

func convertToAscii(r io.Reader) io.Reader {
	s := bufio.NewScanner(r)
	buf := new(bytes.Buffer)
	for s.Scan() {
		buf.WriteString(s.Text())
		buf.WriteString("\r\n")
	}
	return buf
}
