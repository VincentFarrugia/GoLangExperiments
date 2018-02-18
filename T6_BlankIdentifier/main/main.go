package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// Using the blank identifier we mean that we
	// do not care about a particular return type.
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

	// This would cause Go to complain because
	// we would be declaring and assigninng the 'errorData' variable but not using it.
	//res, errorData := http.Get("http://www.google.com")
}
