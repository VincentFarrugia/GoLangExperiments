package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

//////////////////////////////
// LOGIN CONSTANTS
//////////////////////////////

const cLoginFormValueUsername = "username"
const cLoginFormValuePassword = "psw"
const cLoginCookieKeySessionID = "sessionID"

//////////////////////////////
// LOGIN STRUCTS
//////////////////////////////

type loginTemplateInput struct {
	Username    string
	Password    string
	LoginFailed bool
}

//////////////////////////////
// LOGIN ENDPOINT HANDLERS
//////////////////////////////

func loginEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		loginGETEndpoint(w, req)
	case http.MethodPost:
		loginPOSTEndpoint(w, req)
	default:
		unknownErrorEndpoint(w, req)
	}
}

func loginGETEndpoint(w http.ResponseWriter, req *http.Request) {
	if !redirectToHomepageIfAlreadyLoggedIn(w, req) {
		// Return the login page.
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
}

func loginPOSTEndpoint(w http.ResponseWriter, req *http.Request) {
	if !redirectToHomepageIfAlreadyLoggedIn(w, req) {
		// Attempt to login.
		frmUsername := req.FormValue(cLoginFormValueUsername)
		frmPassword := req.FormValue(cLoginFormValuePassword)
		if login(frmUsername, frmPassword, w, req) {
			http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		} else {
			templateInput := loginTemplateInput{
				Username:    "",
				Password:    "",
				LoginFailed: true,
			}
			tpl.ExecuteTemplate(w, "login.gohtml", templateInput)
		}

	}
}

//////////////////////////////
// LOGIN HELPERS
//////////////////////////////

func login(userID string, shadow string, w http.ResponseWriter, req *http.Request) bool {
	ptReqUser := users.GetUser(userID)
	if ptReqUser != nil {

		err := bcrypt.CompareHashAndPassword([]byte(ptReqUser.Shadow), []byte(shadow))
		if err == nil {
			// This is a valid user.
			sessionUUID, err := uuid.NewV4()
			if err != nil {
				// TODO: handle error.
			}
			nwSession := &Session{
				SessionID: sessionUUID.String(),
				UserID:    ptReqUser.UserID,
			}
			sessions.SetSession(nwSession)

			ptSessionCookie := &http.Cookie{
				Name:  cLoginCookieKeySessionID,
				Value: nwSession.SessionID,
			}
			http.SetCookie(w, ptSessionCookie)
			req.AddCookie(ptSessionCookie)

			sessions.SaveToCSVFileIfDirty(cSessionTableFile)
			return true
		}
	}
	return false
}

func isAlreadyLoggedIn(req *http.Request) bool {
	retFlag := false
	sessionID := getSessionIDFromRequest(req)
	if sessions.HasEntry(sessionID) {
		storedUserID := sessions.GetSession(sessionID).UserID
		if users.HasEntry(storedUserID) {
			return true
		}
	}
	return retFlag
}

//////////////////////////////
