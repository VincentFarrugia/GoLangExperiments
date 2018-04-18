/*
Server program which checks for a sessionUUID stored within the
client data in cookie form. A new sessionUUID is created if the
client does not return one.
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

const cParamKeySessionID = "session-id"

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {

	sessionUUIDCookie, err := req.Cookie(cParamKeySessionID)
	if err != nil {
		sessionUUID, err := uuid.NewV4()
		if err == nil {
			sessionUUIDCookie = &http.Cookie{
				Name:     cParamKeySessionID,
				Value:    sessionUUID.String(),
				HttpOnly: true,
			}
			http.SetCookie(w, sessionUUIDCookie)
		}
	}
	fmt.Printf("Server got connection from client with session-id: '%s'\n", sessionUUIDCookie.String())
}
