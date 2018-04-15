/*
Description:
Examples of writing, reading and deleting cookie data.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/writeToClient", writeToClientEndpoint)
	http.HandleFunc("/readFromClient", readFromClientEndpoint)
	http.HandleFunc("/deletePartialFromClient", deletePartialFromClientEndpoint)
	http.HandleFunc("/deleteAllFromClient", deleteAllFromClientEndpoint)
	http.ListenAndServe(":8080", nil)
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {

}

func writeToClientEndpoint(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "ItemA",
		Value: "This is value A",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "ItemB",
		Value: "This is value B",
	})
}

func readFromClientEndpoint(w http.ResponseWriter, req *http.Request) {
	cookieKeyList := []string{"ItemA", "ItemB"}
	io.WriteString(w, "Cookie Data:\n")
	for _, itemKey := range cookieKeyList {
		ptCookie, err := req.Cookie(itemKey)
		reqValue := "NULL"
		if (ptCookie != nil) && (err == nil) {
			reqValue = ptCookie.Value
		}
		io.WriteString(w, fmt.Sprintf("%s:%s\n", itemKey, reqValue))
	}

	// NOTE: We could use err == http.ErrNoCookie to be more precise than just err != nil.
}

func deletePartialFromClientEndpoint(w http.ResponseWriter, req *http.Request) {
	// Eg. deleting cookie data for key "ItemB"
	// Effectively, setting the data MaxAge to 0 or negative seconds
	// will delete it on the client machine.
	// (using Expires instead of MaxAge is not recommended as it is deprecated)
	deleteCookie(w, "ItemB")
}

func deleteAllFromClientEndpoint(w http.ResponseWriter, req *http.Request) {
	deleteCookie(w, "ItemA")
	deleteCookie(w, "ItemB")
}

/////////////////////////////////////////
// HELPER FUNCTIONS:
/////////////////////////////////////////

func deleteCookie(w http.ResponseWriter, cookieKey string) {
	http.SetCookie(w, &http.Cookie{
		Name:   cookieKey,
		Value:  "",
		MaxAge: -1,
		//Expires: time.Unix(0, 0),   // This is deprecated. Use MaxAge instead.
	})
}

func deleteCookies(w http.ResponseWriter, cookieKeyList []string) {
	for _, cookieKey := range cookieKeyList {
		deleteCookie(w, cookieKey)
	}
}

/////////////////////////////////////////
