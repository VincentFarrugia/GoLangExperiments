// Basic example of net/http package usage of ListenAndServe.
// No muxes, just a basic handler for this example.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := myDefaultHTTPHandler{}
	http.ListenAndServe(":8080", h)
}

////////////////////////////////////////////////
// HANDLERS
////////////////////////////////////////////////

type myDefaultHTTPHandler struct {
}

func (h myDefaultHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You have accessed route: '%s'", req.RequestURI)
}

////////////////////////////////////////////////
