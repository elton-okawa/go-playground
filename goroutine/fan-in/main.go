package main

import (
	"fmt"
	"math/rand"
	"time"
)

func message(value string) <-chan string {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		c <- value
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case v := <-input1:
				c <- v
			case v := <-input2:
				c <- v
			}
		}
	}()

	return c
}

func main() {
	c := fanIn(message("Hello"), message("World"))

	// there is no guarantee to be ordered
	fmt.Println(<-c)
	fmt.Println(<-c)
}
