package main

import "net/http"

/////////////////////////////////
// WELCOME CONSTANTS
/////////////////////////////////

const cWelcomeCookieSessionID = "sessionID"

/////////////////////////////////
// WELCOME STRUCTS
/////////////////////////////////

type welcomeTemplateInput struct {
	Username string
	Role     string
}

/////////////////////////////////
// WELCOME ENDPOINT HANDLERS
/////////////////////////////////

// TODO: Add security.
func welcomeEndpoint(w http.ResponseWriter, req *http.Request) {
	bSuccess := false
	sessionCookie, err := req.Cookie(cWelcomeCookieSessionID)
	if err == nil {
		sessionID := sessionCookie.Value
		if sessions.HasEntry(sessionID) {
			reqUserID := sessions.GetSession(sessionID).UserID
			ptReqUser := users.GetUser(reqUserID)
			if ptReqUser != nil {
				templateInput := welcomeTemplateInput{
					Username: ptReqUser.UserID,
					Role:     ptReqUser.Role,
				}
				tpl.ExecuteTemplate(w, "welcome.gohtml", templateInput)
				bSuccess = true
			}
		}
	}

	if !bSuccess {
		unknownErrorEndpoint(w, req)
	}
}

/////////////////////////////////
