/*
Description:
Example using HTTP redirects.
HTTP 301 - Moved permanently.
HTTP 303 - See other.
HTTP 307 - Temporary redirect.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/itemA", itemAEndpoint)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", rootEndpoint)
	http.ListenAndServe(":8080", nil)
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("The is the response for requesting on endpoint: '%s'", req.RequestURI))
}

func itemAEndpoint(w http.ResponseWriter, req *http.Request) {

	// Basic example by setting the header yourself.
	//w.Header().Set("Location", "/itemB")
	//w.WriteHeader(http.StatusTemporaryRedirect)

	// This is an example of doing similar things using the Redirect function in the net/http package.
	http.Redirect(w, req, "/itemB", http.StatusTemporaryRedirect)
}
