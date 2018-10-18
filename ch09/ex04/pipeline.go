package main

import "flag"
import "log"

func main() {
	length := flag.Int("length", 10, "length of pipeline")
	flag.Parse()

	start := make(chan string)
	channels := []chan string{start}

	for i := 0; i < *length; i++ {
		out := make(chan string)
		go pass(out, channels[len(channels)-1])
		channels = append(channels, out)
	}

	log.Println("start")
	start <- "Hello, World!"
	end := <-channels[len(channels)-1]
	log.Printf("end: %q\n", end)
}

func pass(out chan<- string, in <-chan string) {
	recieved := <-in
	//log.Printf("recieved: %q\n", recieved)
	out <- recieved
}
