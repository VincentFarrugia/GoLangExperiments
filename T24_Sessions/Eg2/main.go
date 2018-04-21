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
var users UserTable
var sessions SessionTable

//////////////////////
// CONSTANTS
//////////////////////

const cCookieKeySessionID = "sessionID"
const cTemplatesDir = "Templates"
const cDataDir = "Data"
const cUserTableFile = cDataDir + "/userTable.csv"
const cSessionTableFile = cDataDir + "/sessionTable.csv"

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
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

///////////////////////////
// HTTP ENDPOINT HANDLERS
///////////////////////////

func rootEndpoint(w http.ResponseWriter, req *http.Request) {

	var ptSessionData *Session
	var ptUserData *User

	sessionID := getSessionIDFromRequest(req)
	ptSessionData = sessions.GetSession(sessionID)
	if ptSessionData == nil {
		ptUserData = users.AddNewBlankUser()
		ptSessionData = sessions.AddNewBlankSessionWithUser(ptUserData.UserID)
	}
	ptUserData = users.GetUser(ptSessionData.UserID)
	if ptUserData == nil {
		ptUserData = users.AddNewBlankUser()
		ptSessionData.AssignUserToSession(ptUserData)
		sessions.SetSession(ptSessionData)
	}
	users.SaveToCSVFileIfDirty(cUserTableFile)
	sessions.SaveToCSVFileIfDirty(cSessionTableFile)

	http.SetCookie(w, &http.Cookie{
		Name:  cCookieKeySessionID,
		Value: ptSessionData.SessionID,
	})

	tpl.ExecuteTemplate(w, "homepage.gohtml", *ptSessionData)
}

/////////////////////////
// HELPER FUNCTIONS
/////////////////////////

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
