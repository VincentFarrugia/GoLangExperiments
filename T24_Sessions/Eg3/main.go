//////////////////////////////////////////////////////////////
// Basic experiment for user registration / login / logout.
//////////////////////////////////////////////////////////////

package main

import (
	"html/template"
	"net/http"
)

//////////////////////
// VARS
//////////////////////

var tpl *template.Template
var users UserTable
var sessions SessionTable

//////////////////////
// CONSTANTS
//////////////////////

const cTemplatesDir = "Templates"
const cDataDir = "Data"
const cUserTableFile = cDataDir + "/userTable.csv"
const cSessionTableFile = cDataDir + "/sessionTable.csv"

const cCookieKeySessionID = "sessionID"

/////////////////////////
// GO PROGRAM LIFECYCLE
/////////////////////////

func init() {
	tpl = template.Must(template.ParseGlob(cTemplatesDir + "/*.gohtml"))
	users.InitFromCSVFile("UserTable", cUserTableFile, &users)
	sessions.InitFromCSVFile("SessionTable", cSessionTableFile, &sessions)
}

func main() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/register", registerUserEndpoint)
	http.HandleFunc("/login", loginEndpoint)
	http.HandleFunc("/logout", logoutEndpoint)
	http.HandleFunc("/welcome", welcomeEndpoint)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

///////////////////////////
// HTTP ENDPOINT HANDLERS
///////////////////////////

func rootEndpoint(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func unknownErrorEndpoint(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Resource not found", http.StatusNotFound)
}

/////////////////////////
// HELPER FUNCTIONS
/////////////////////////

func redirectToHomepageIfAlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	if isAlreadyLoggedIn(req) {
		// Redirect to welcome page.
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return true
	}
	return false
}

func getSessionIDFromRequest(req *http.Request) string {
	retSessionID := ""
	if req != nil {
		sessionIDCookie, err := req.Cookie(cCookieKeySessionID)
		if err == nil {
			retSessionID = sessionIDCookie.Value
		}
	}
	return retSessionID
}

/////////////////////////
