////////////////////////////////////////////////////////////////////////////////
// WORK IN PROGRESS
// Experimental code for an http server without using the net.http package.
// This is just as a fun challenge and golang practice.
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"net"

	"github.com/VincentFarrugia/GoLangExperiments/T19_TCP/Eg4/httpUtils"
)

const cWebContentRelativeRoot = "WebContent"

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
	request := httpUtils.ParseHTTPRequest(clientConn)
	//fmt.Println(request)

	response := httpUtils.HTTPResponse{}
	response.Constructor()
	switch request.Method {
	case httpUtils.CHTTPMethodGet:
		{
			response = getResource(request)
		}
	}
	httpUtils.SendHTTPResponse(clientConn, response)
}

//////////////////////////////////////////
// SERVER HTTP MUX ENDPOINT HANDLERS:
//////////////////////////////////////////

func getResource(request httpUtils.HTTPRequest) httpUtils.HTTPResponse {

	// Default response set to server error.
	response := httpUtils.HTTPResponse{}
	response.ConstructorWithStatusLine("HTTP/1.1", 500, "Server Error")

	contentStr := ""
	bFoundResource := false

	if request.Method == httpUtils.CHTTPMethodGet {
		if request.RequestURI != "" {

			relativeURI := cWebContentRelativeRoot
			if request.RequestURI[0] == '/' {
				relativeURI += request.RequestURI[1:]
			}

			if relativeURI == cWebContentRelativeRoot {
				relativeURI += "/index.html"
			}

			contentStr, bFoundResource = httpUtils.GetResourceFileContents(relativeURI)
		}
	}

	if bFoundResource {
		response.SetBody(contentStr)
	} else {
		response.SetStatusCode(404, "Could not find resource")
	}

	return response
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
