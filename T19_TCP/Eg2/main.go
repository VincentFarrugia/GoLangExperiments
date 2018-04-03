package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
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
	defer clientConn.Close()

	err := clientConn.SetDeadline(time.Now().Add(20 * time.Second))
	panicIfErr(err)
	scanner := bufio.NewScanner(clientConn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
