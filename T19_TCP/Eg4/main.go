////////////////////////////////////////////////////////////////////////////////
// WORK IN PROGRESS
// Experimental code for an http server without using the net.http package.
// This is just as a fun challenge and golang practice.
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"net"
)

const (
	//httpMethodOptions = "OPTIONS"
	httpMethodGet = "GET"
	/*httpMethodHead    = "HEAD"
	httpMethodPost    = "POST"
	httpMethodPut     = "PUT"
	httpMethodDelete = "DELETE"
	httpMethodTrace = "TRACE"
	httpMethodConnect = "CONNECT"*/
)

type httpRequest struct {
	// Request-Line
	method      string
	requestURI  string
	httpVersion string
	// Headers
	// TODO: expand this into more detail.
	headers map[string]string
	// Message Body
	body string
}

type httpResponse struct {
}

func main() {
	fmt.Println("********************************")
	fmt.Println("HTTP Server Initialised.")
	fmt.Println("********************************")
	listener, err := net.Listen("tcp", "localhost:8080")
	panicIfErr(err)
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting Client connection. Cancelling request.")
			clientConn.Close()
		} else {
			go handleClient(clientConn)
		}
	}
}

//////////////////////////////////////////
// GOROUTINES:
//////////////////////////////////////////

func handleClient(clientConn net.Conn) {
	defer clientConn.Close()
	/*scanner := bufio.NewScanner(clientConn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	fmt.Println("End Of File")*/
}

//////////////////////////////////////////
// HELPER FUNCTIONS - HTTP:
//////////////////////////////////////////

func parseHTTPRequest(rawRequestData string) httpRequest {
	retData := httpRequest{}
	return retData
}

//////////////////////////////////////////
// HELPER FUNCTIONS - BASIC:
//////////////////////////////////////////

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

//////////////////////////////////////////
