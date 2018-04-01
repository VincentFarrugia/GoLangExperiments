/*
Basic TCP server example 1.
*/

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	panicIfError(err)
	defer listener.Close()

	// Accept and handle clients indefinitely.
	for {

		// Accept
		connToClient, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		// Do something with client connection and then close connection.
		fmt.Fprintln(connToClient, "Hi there Client! This is the Server speaking!")
		connToClient.Close()
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
