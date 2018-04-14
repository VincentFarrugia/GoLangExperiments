package main

import (
	"io"
	"net/http"
)

// http.FileServer example.
func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/mainPage", mainPageHandleFunc)
	http.ListenAndServe(":8181", nil)
}

func mainPageHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="resources/homeImg3.png">`)
}
