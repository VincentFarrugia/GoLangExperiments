package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", rootEndpointHandler)
	http.HandleFunc("/myImageB", myImageBEndpointHandler)
	http.ListenAndServe(":8181", nil)
}

// Example of serving a file which we have locally on our server.
func rootEndpointHandler(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("homeImg.png")
	panicIfErr(err)
	defer file.Close()
	io.Copy(w, file)
}

// Example of including an html element which indirectly causes
// the client browser to ask a third-party server for a resource,
// in this case an image.
func myImageBEndpointHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html><img src=\"https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Golang.png/220px-Golang.png\" alt=\"GitHub image\"/></html>")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
