package main

import "net/http"

/////////////////////////////////
// LOGOUT CONSTANTS
/////////////////////////////////

const cLogoutCookieKeySessionID = "sessionID"

/////////////////////////////////
// LOGOUT ENDPOINT HANDLERS
/////////////////////////////////

func logoutEndpoint(w http.ResponseWriter, req *http.Request) {

	sessionIDCookie, err := req.Cookie(cLogoutCookieKeySessionID)
	if err == nil {
		sessionID := sessionIDCookie.Value
		sessions.RemoveEntry(sessionID)
		sessions.SaveToCSVFileIfDirty(cSessionTableFile)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cLogoutCookieKeySessionID,
		Value:  "",
		MaxAge: -1,
	})

	tpl.ExecuteTemplate(w, "goodbye.gohtml", nil)

}

/////////////////////////////////
