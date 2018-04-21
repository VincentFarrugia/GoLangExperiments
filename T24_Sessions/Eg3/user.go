package main

import (
	"fmt"
)

//////////////////////////
// USER
//////////////////////////

// User represents information for a single user
// registered with the server.
type User struct {
	UserID    string
	FirstName string
	Surname   string
	Email     string
	Shadow    string
	Role      string
}

// IsUserRole returns true if the passed in role string
// matches the role for the given user.
func (usr *User) IsUserRole(role string) bool {
	return usr.Role == role
}

// InitBlank (DataTableEntry interface)
// Create a blank User.
func (usr *User) InitBlank() {
	if usr != nil {
		usr.UserID = ""
		usr.FirstName = ""
		usr.Surname = ""
		usr.Email = ""
		usr.Shadow = ""
		usr.Role = ""
	}
}

// GetPrimaryKey (DataTableEntry interface)
// Returns the User's UUID.
func (usr *User) GetPrimaryKey() string {
	if usr != nil {
		return usr.UserID
	}
	return ""
}

// SetFromCSVLine (DataTableEntry interface)
// Sets a User from a CSV line string.
func (usr *User) SetFromCSVLine(csvLine []string) {
	if usr != nil {
		if len(csvLine) == 6 {
			usr.UserID = csvLine[0]
			usr.FirstName = csvLine[1]
			usr.Surname = csvLine[2]
			usr.Email = csvLine[3]
			usr.Shadow = csvLine[4]
			usr.Role = csvLine[5]
		}
	}
}

// ToCSVLine (DataTableEntry interface)
// Returns the User data as a CSV line string.
func (usr *User) ToCSVLine() string {
	if usr != nil {
		return (fmt.Sprintf("%s,%s,%s,%s,%s,%s", usr.UserID, usr.FirstName, usr.Surname, usr.Email, usr.Shadow, usr.Role))
	}
	return ""
}

//////////////////////////
