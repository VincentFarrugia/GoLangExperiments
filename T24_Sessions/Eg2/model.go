///////////////////////////////////////////////////////////////////////////
//	Work in progress.
//	Example of session creation / user creation from scratch.
///////////////////////////////////////////////////////////////////////////

package main

import uuid "github.com/satori/go.uuid"

//////////////////////////
// USERS AND USER-TABLE
//////////////////////////

type user struct {
	UserID    string
	FirstName string
	Surname   string
	Email     string
}

type userTable map[string]user

func newUserTable() userTable {
	return userTable(make(map[string]user))
}

func (ut *userTable) userExists(userID string) bool {
	if usr, bOk := (*ut)[userID]; bOk {
		return true
	}
	return false
}

func (ut *userTable) getUser(userID string) *user {
	if usr, bOk := (*ut)[userID]; bOk {
		return &usr
	}
	return nil
}

func (ut *userTable) setUser(userData user) {
	if userData.UserID != "" {
		(*ut)[userData.UserID] = userData
	}
}

func (ut *userTable) addNewBlankUser() *user {
	uuid, _ := uuid.NewV4()
	nwUser := user{}
	nwUser.UserID = uuid.String()
	(*ut)[nwUser.UserID] = nwUser
	return ut.getUser(nwUser.UserID)
}

///////////////////////////////
// SESSION AND SESSION-TABLE
///////////////////////////////

type session struct {
	SessionID string
	UserID    string
}

func (s *session) assignUserToSession(usr *user) bool {
	if usr != nil && usr.UserID != "" {
		s.UserID = usr.UserID
		return true
	}
	return false
}

type sessionTable map[string]session

func newSessionTable() sessionTable {
	return sessionTable(make(map[string]session))
}

func (st *sessionTable) sessionExists(sessionID string) bool {
	if s, bOk := (*st)[sessionID]; bOk {
		return true
	}
	return false
}

func (st *sessionTable) getSession(sessionID string) *session {
	if s, bOk := (*st)[sessionID]; bOk {
		return &s
	}
	return nil
}

func (st *sessionTable) setSession(sessionData session) {
	if sessionData.SessionID != "" {
		(*st)[sessionData.SessionID] = sessionData
	}
}

func (st *sessionTable) addNewBlankSession() *session {
	uuid, _ := uuid.NewV4()
	nwSession := session{}
	nwSession.SessionID = uuid.String()
	(*st)[nwSession.SessionID] = nwSession
	return st.getSession(nwSession.SessionID)
}

///////////////////////////////
