package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.ListenAndServe(":8080", nil)
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world! DKR Container")
}
