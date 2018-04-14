package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", rootEndpointHandler)
	http.HandleFunc("/homeImg2.png", homeImageEndpointHandler)
	http.HandleFunc("/homeImg3.png", homeImage3EndpointHandler)
	http.ListenAndServe(":8181", nil)
}

func rootEndpointHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/homeImg2.png"> <img src="/homeImg3.png"`)
}

// Example using http.ServeContent.
func homeImageEndpointHandler(w http.ResponseWriter, req *http.Request) {

	file, err := os.Open("homeImg2.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, req, file.Name(), fileInfo.ModTime(), file)
}

// Example using http.ServeFile.
func homeImage3EndpointHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "homeImg3.png")
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
