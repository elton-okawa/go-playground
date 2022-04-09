package main

import (
	"fmt"
	"strconv"
	"time"
)

func producer(ms time.Duration) <-chan string {
	c := make(chan string, 3)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(ms * time.Millisecond)
			v := strconv.Itoa(i)
			fmt.Printf("[%s] Sending: %s\n", time.Now().Format(time.StampMilli), v)
			c <- v
		}
	}()

	return c
}

func consumer(channel <-chan string, ms time.Duration) {
	for i := 0; i < 10; i++ {
		time.Sleep(ms * time.Millisecond)
		v := <-channel
		fmt.Printf("[%s] Received: %s\n", time.Now().Format(time.StampMilli), v)
	}
}

// Setup
// - producer has a buffered channel of size 3
// - producer is 5 times faster than the consumer
// From the output we can observe 3 phases:
// - producer runs until fills the buffer and become blocked
// - producer and consumer alternate between producing/consuming value
// - producer finishes it loops and consumer consumes the remaining values on buffer
func main() {
	c := producer(100)
	consumer(c, 500)
}
