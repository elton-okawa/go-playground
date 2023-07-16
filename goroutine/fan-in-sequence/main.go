package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func loopMessage(value string, long bool) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	current := long
	go func() {
		for i := 0; i < 5; i++ {
			c <- Message{fmt.Sprintf("%s: %d (long: %t)", value, i, current), waitForIt}
			if current {
				time.Sleep(time.Duration(10000) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(3000) * time.Millisecond)
			}
			current = !current
			<-waitForIt
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
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
	c := fanIn(loopMessage("Hello", true), loopMessage("World", false))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Printf("[%s] %s\n", time.Now().Format(time.StampMilli), msg1.str)
		msg2 := <-c
		fmt.Printf("[%s] %s\n", time.Now().Format(time.StampMilli), msg2.str)
		fmt.Println()

		// The slowest one dictates
		msg1.wait <- true
		msg2.wait <- true
	}
}
