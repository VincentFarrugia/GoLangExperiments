package main

import (
	"fmt"
	"time"
)

///////////////////////////////
// SESSION
///////////////////////////////

const cSessionLastActivityTimeFormat = "2006-01-02T15:04:05.000Z"
const cSessionExpireTimeSpan = (time.Second * 60 * 5)

// Session represents information about a single client-server session.
// It is associated with one UserID.
type Session struct {
	SessionID    string
	UserID       string
	LastActivity time.Time
}

// InitBlank (DataTableEntry interface)
// Create blank Session info.
func (s *Session) InitBlank() {
	if s != nil {
		s.SessionID = ""
		s.UserID = ""
		s.LastActivity = time.Now()
	}
}

// GetPrimaryKey (DataTableEntry interface)
// Returns the Session's UUID.
func (s *Session) GetPrimaryKey() string {
	if s != nil {
		return s.SessionID
	}
	return ""
}

// SetFromCSVLine (DataTableEntry interface)
// Sets a Session from a CSV line string.
func (s *Session) SetFromCSVLine(csvLine []string) {
	if s != nil {
		if len(csvLine) == 3 {
			s.SessionID = csvLine[0]
			s.UserID = csvLine[1]
			t, err := time.Parse(cSessionLastActivityTimeFormat, csvLine[2])
			if err != nil {
				s.LastActivity = time.Now()
			} else {
				s.LastActivity = t
			}
		}
	}
}

// ToCSVLine (DataTableEntry interface)
// Returns the Session data as a CSV line string.
func (s *Session) ToCSVLine() string {
	if s != nil {
		return (fmt.Sprintf("%s,%s,%s", s.SessionID, s.UserID, s.LastActivity.Format(cSessionLastActivityTimeFormat)))
	}
	return ""
}

// AssignUserToSession associates a User with a Session
func (s *Session) AssignUserToSession(usr *User) bool {
	if usr != nil && usr.UserID != "" {
		s.UserID = usr.UserID
		return true
	}
	return false
}

///////////////////////////////
