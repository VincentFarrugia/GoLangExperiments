package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootEndpointHandleFunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func rootEndpointHandleFunc(w http.ResponseWriter, req *http.Request) {
	ipsumVal := req.FormValue("ipsum")
	if ipsumVal == "" {
		io.WriteString(w, "Server did not receive a value for ipsum!")
	} else {
		io.WriteString(w, fmt.Sprintf("Server received ipsum value: '%s'", ipsumVal))
	}
}
