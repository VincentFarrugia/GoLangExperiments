package main

import (
	"net/http"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

/////////////////////////////////
// REGISTER CONSTANTS
/////////////////////////////////

const cRegisterFormValueUsername = "username"
const cRegisterFormValuePassword = "psw"
const cRegisterFormValuePasswordConfirm = "psw2"
const cRegisterFormValueFirstname = "firstname"
const cRegisterFormValueSurname = "surname"
const cRegisterFormValueEmail = "email"
const cRegisterFormValueSessionID = "sessionID"
const cRegisterFormValueRole = "role"

/////////////////////////////////
// REGISTER STRUCTS
/////////////////////////////////

type registerTemplateInput struct {
	Username          string
	Password          string
	PasswordConfirm   string
	Firstname         string
	Surname           string
	Email             string
	Role              string
	UsernameAvailable bool
	PasswordsMatched  bool
	RegisterFailed    bool
}

/////////////////////////////////
// REGISTER ENDPOINT HANDLERS
/////////////////////////////////

func registerUserEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		registerUserGETEndpoint(w, req)
	case http.MethodPost:
		registerUserPOSTEndpoint(w, req)
	default:
		unknownErrorEndpoint(w, req)
	}
}

func registerUserGETEndpoint(w http.ResponseWriter, req *http.Request) {
	// Return the register page.
	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}

func registerUserPOSTEndpoint(w http.ResponseWriter, req *http.Request) {
	// Attempt to register user.
	frmUsername := req.FormValue(cRegisterFormValueUsername)
	frmPassword := req.FormValue(cRegisterFormValuePassword)
	frmPassword2 := req.FormValue(cRegisterFormValuePasswordConfirm)
	frmFirstname := req.FormValue(cRegisterFormValueFirstname)
	frmSurname := req.FormValue(cRegisterFormValueSurname)
	frmEmail := req.FormValue(cRegisterFormValueEmail)
	frmRole := req.FormValue(cRegisterFormValueRole)
	bUsernameAvailable := (users.HasEntry(frmUsername) == false)
	bPasswordsMatched := (frmPassword == frmPassword2)

	if bUsernameAvailable && bPasswordsMatched {

		// Valid information given.
		// Register the new user.

		bs, err := bcrypt.GenerateFromPassword([]byte(frmPassword), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		nwUser := &User{
			UserID:    frmUsername,
			FirstName: frmFirstname,
			Surname:   frmSurname,
			Email:     frmEmail,
			Shadow:    string(bs),
			Role:      frmRole,
		}
		users.SetUser(nwUser)

		sessionUUID, err := uuid.NewV4()
		if err != nil {
			// TODO: handle error.
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		nwSession := &Session{
			SessionID: sessionUUID.String(),
			UserID:    nwUser.UserID,
		}
		sessions.SetSession(nwSession)

		ptSessionCookie := &http.Cookie{
			Name:  cRegisterFormValueSessionID,
			Value: nwSession.SessionID,
		}
		http.SetCookie(w, ptSessionCookie)
		req.AddCookie(ptSessionCookie)

		users.SaveToCSVFileIfDirty(cUserTableFile)
		sessions.SaveToCSVFileIfDirty(cSessionTableFile)

		// Redirect to the welcome page.
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)

	} else {
		templateInput := registerTemplateInput{
			Username:          frmUsername,
			Password:          frmPassword,
			PasswordConfirm:   frmPassword2,
			Firstname:         frmFirstname,
			Surname:           frmSurname,
			Email:             frmEmail,
			Role:              frmRole,
			UsernameAvailable: bUsernameAvailable,
			PasswordsMatched:  bPasswordsMatched,
			RegisterFailed:    true,
		}
		tpl.ExecuteTemplate(w, "register.gohtml", templateInput)
	}
}

/////////////////////////////////
