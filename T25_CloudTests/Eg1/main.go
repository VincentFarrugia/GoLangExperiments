package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootEndpoint)
	err := http.ListenAndServe(":80", nil)
	log.Fatalln(err.Error())
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World AWS!")
}
