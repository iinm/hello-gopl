package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
	counter := make(chan struct{})

	go pingpong(ping, pong, counter)
	go pingpong(pong, ping, counter)
	ping <- struct{}{}

	var count int
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("ping-pong count: %d / sec\n", count)
			count = 0
		case <-counter:
			count++
		}
	}
	//ticker.Stop()
}

func pingpong(in <-chan struct{}, out chan<- struct{}, counter chan<- struct{}) {
	for range in {
		out <- struct{}{}
		counter <- struct{}{}
	}
}
