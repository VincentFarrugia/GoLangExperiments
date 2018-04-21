package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

///////////////////////////////
// SESSION TABLE
///////////////////////////////

// SessionTable stores session entries.
type SessionTable struct {
	DataTable
}

// GetSession is a helper function ontop of GetEntry
// which returns a pointer to a Session
func (st *SessionTable) GetSession(sessionID string) *Session {
	if st.HasEntry(sessionID) {
		dte := st.GetEntry(sessionID)
		return dte.(*Session)
	}
	return nil
}

// SetSession is a helper function ontop of SetEntry
func (st *SessionTable) SetSession(s *Session) {
	st.SetEntry(s.SessionID, s)
}

// AddNewBlankSessionWithUser creates a new session with a newly generated UUID.
// The passed in UserID is associated with the new session.
// If the userID is empty (empty string), the session is not created and nil is returned.
func (st *SessionTable) AddNewBlankSessionWithUser(userID string) *Session {
	if userID == "" {
		return nil
	}
	uuid, _ := uuid.NewV4()
	nwSession := Session{}
	nwSession.SessionID = uuid.String()
	nwSession.UserID = userID
	st.SetEntry(nwSession.SessionID, &nwSession)
	return st.GetSession(nwSession.SessionID)
}

// CleanSessions deletes any sessions which have remained for too long
// depending on the session expire timespan.
func (st *SessionTable) CleanSessions() {
	itemsToRemove := make([]string, 0)
	for sessionID := range st.Rows {
		ptReqSessionData := st.GetSession(sessionID)
		if ptReqSessionData != nil {
			if time.Now().Sub(ptReqSessionData.LastActivity) > (cSessionExpireTimeSpan) {
				itemsToRemove = append(itemsToRemove, sessionID)
			}
		}
	}
	for _, sessionID := range itemsToRemove {
		st.RemoveEntry(sessionID)
		st.SetDirty()
	}
}

// CreateBlankEntry (IBlankTableEntryCreator interface)
func (st *SessionTable) CreateBlankEntry() DataTableEntry {
	return &Session{}
}

///////////////////////////////
