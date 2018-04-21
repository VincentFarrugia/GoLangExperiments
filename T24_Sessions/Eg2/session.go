package main

import (
	"fmt"
)

///////////////////////////////
// SESSION
///////////////////////////////

// Session represents information about a single client-server session.
// It is associated with one UserID.
type Session struct {
	SessionID string
	UserID    string
}

// InitBlank (DataTableEntry interface)
// Create blank Session info.
func (s *Session) InitBlank() {
	if s != nil {
		s.SessionID = ""
		s.UserID = ""
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
		if len(csvLine) == 2 {
			s.SessionID = csvLine[0]
			s.UserID = csvLine[1]
		}
	}
}

// ToCSVLine (DataTableEntry interface)
// Returns the Session data as a CSV line string.
func (s *Session) ToCSVLine() string {
	if s != nil {
		return (fmt.Sprintf("%s,%s", s.SessionID, s.UserID))
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
