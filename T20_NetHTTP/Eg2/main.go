//
// Example for using the default golang http package servermux.
//

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Using the default golang http package servermux.
	http.HandleFunc("/home", homeEndpoint)
	http.HandleFunc("/testing/", testingEndpoint)
	http.ListenAndServe(":8080", nil)

	// Note: The setup above will cause the following to happen:
	// Requests on "/home" - passed to the homeEndpoint handler func.
	// Requests on "/home/anything/something" - not handled.
	// Requests on "/testing" - passed to the testingEndpoint handler func.
	// Requests on "/testing/anything/whatever" - passed to the testingEndpoint handler func.
}

////////////////////////////////
// ENDPOINT HANDLER FUNCTIONS:
////////////////////////////////

func homeEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the 'home' endpoint handler!")
}

func testingEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "This is the 'testing' endpoint handler!")
}

////////////////////////////////
