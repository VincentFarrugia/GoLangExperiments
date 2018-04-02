package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	panicIfErr(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		panicIfErr(err)
		go handleClient(conn)
	}
}

func handleClient(clientConn net.Conn) {
	fmt.Fprintln(clientConn, "Hello client!")
	fmt.Println("Server goroutine connected to client.")
	fmt.Println("Server goroutine now reading client text:")
	// TODO: Implement closing of stream.
	scanner := bufio.NewScanner(clientConn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	clientConn.Close()
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
