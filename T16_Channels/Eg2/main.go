package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// This essentially seals of the channel.
		// (Closes the pipe (in C lang terms))
		close(c)
	}()

	// Will keep reading in data from channel
	// into variable 'n' until we notice that
	// the channel has been empty and closed.
	for n := range c {
		fmt.Println(n)
	}
}
