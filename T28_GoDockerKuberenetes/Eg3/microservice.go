package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	bookAPI "github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg2/api"
)

var bookDatabase = map[string]bookAPI.Book{
	"A_ISBN": bookAPI.Book{Name: "A_Name", Author: "A_Author", ISBN: "A_ISBN"},
	"B_ISBN": bookAPI.Book{Name: "B_Name", Author: "B_Author", ISBN: "B_ISBN"},
	"C_ISBN": bookAPI.Book{Name: "C_Name", Author: "C_Author", ISBN: "C_ISBN"},
}

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/echo", echoEndpoint)
	http.HandleFunc("/api/books", apiBooksEndpoint)
	http.HandleFunc("/api/books/", apiBooksModEndpoint)
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
	switch req.Method {
	case http.MethodGet:
		writeJSON(w, bookDatabase)
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := bookAPI.FromJSON(body)
		isbn, bCreated := addBookToDatabase(&book)
		if bCreated {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported API Request METHOD")
	}
}

func apiBooksModEndpoint(w http.ResponseWriter, req *http.Request) {
	// First, get ISBN as last part of the URL.
	isbn := req.URL.Path[len("/api/books/"):]

	switch req.Method {
	case http.MethodGet:
		book, bFound := getBookFromDatabase(isbn)
		if bFound {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := bookAPI.FromJSON(body)
		bExists := updateBookInDatabase(&book)
		if bExists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		deleteBookInDatabase(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported request method.")
	}
}

func writeJSON(w http.ResponseWriter, item interface{}) {
	bytes, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}

func addBookToDatabase(b *bookAPI.Book) (string, bool) {
	if b == nil {
		return "", false
	}
	bookDatabase[b.ISBN] = *b
	return b.ISBN, true
}

func updateBookInDatabase(b *bookAPI.Book) bool {
	if b == nil {
		return false
	}

	_, bExists := bookDatabase[b.ISBN]
	if !bExists {
		return false
	} else {
		bookDatabase[b.ISBN] = *b
		return true
	}
}

func deleteBookInDatabase(isbn string) {
	_, bExists := bookDatabase[isbn]
	if !bExists {
		//
	} else {
		delete(bookDatabase, isbn)
	}
}

func getBookFromDatabase(isbn string) (*bookAPI.Book, bool) {
	reqBook, bExists := bookDatabase[isbn]
	if bExists {
		return &reqBook, true
	} else {
		return nil, false
	}
}
