///////////////////////////////////////////////////////////////////////////
//	Work in progress.
//	Example of session creation / user creation from scratch.
///////////////////////////////////////////////////////////////////////////

package main

import (
	"html/template"
	"net/http"
)

//////////////////////
// VARS
//////////////////////

var tpl *template.Template
var users = newUserTable()
var sessions = newSessionTable()

//////////////////////
// CONSTANTS
//////////////////////

const cCookieKeySessionID = "sessionID"

//////////////////////
// FUNCTIONS
//////////////////////

func init() {
	tpl = template.Must(template.ParseGlob("Templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func rootEndpoint(w http.ResponseWriter, req *http.Request) {

	var ptSessionData *session
	var ptUserData *user

	sessionID := getSessionIDFromRequest(req)
	ptSessionData = sessions.getSession(sessionID)
	if ptSessionData == nil {
		ptUserData = users.addNewBlankUser()
		ptSessionData = sessions.addNewBlankSession()
		ptSessionData.UserID = ptUserData.UserID
	}
	if ptUserData == nil {
		ptUserData = users.addNewBlankUser()
	}
	sessionID = ptSessionData.SessionID
	users.setUser(*ptUserData)
	ptSessionData.assignUserToSession(ptUserData)
	sessions.setSession(*ptSessionData)

	tpl.ExecuteTemplate(w, "homepage.gohtml", *ptUserData)
}

func getSessionIDFromRequest(req *http.Request) string {
	retSessionID := ""
	if req != nil {
		sessionIDCookie, err := req.Cookie(cCookieKeySessionID)
		if err == nil {
			retSessionID = sessionIDCookie.String()
		}
	}
	return retSessionID
}
