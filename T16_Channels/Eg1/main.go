package main

import (
	"fmt"
	"time"
)

// Go Proverb: "Don't communicate by sharing memory, share memory by communicating."
//https://go-proverbs.github.io/

func main() {

	// Create a channel (unbuffered).
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			// This will put an int value onto the channel
			// then block until some other go routine
			// reads in that value from the channel.
			c <- i
		}
	}()

	go func() {
		for {
			// If no value currently exists on the channel,
			// this go routine will block.
			// If there is a value, read it.
			fmt.Println("Counter:", <-c)
		}
	}()

	time.Sleep(3 * time.Second)
}
