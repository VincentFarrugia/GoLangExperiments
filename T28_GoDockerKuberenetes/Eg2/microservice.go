package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	bookAPI "github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg2/api"
)

var bookList []bookAPI.Book = []bookAPI.Book{
	bookAPI.Book{
		Name:   "A_Name",
		Author: "A_Author",
		ISBN:   "A_ISBN",
	},
	bookAPI.Book{
		Name:   "B_Name",
		Author: "B_Author",
		ISBN:   "B_ISBN",
	},
	bookAPI.Book{
		Name:   "C_Name",
		Author: "C_Author",
		ISBN:   "C_ISBN",
	},
}

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/echo", echoEndpoint)
	http.HandleFunc("/books", apiBooksEndpoint)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World Yo!")
}

func echoEndpoint(w http.ResponseWriter, req *http.Request) {
	message := req.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func apiBooksEndpoint(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(bookList)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
