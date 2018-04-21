package main

import uuid "github.com/satori/go.uuid"

//////////////////////////
// USER TABLE
//////////////////////////

// UserTable stores user entries.
type UserTable struct {
	DataTable
}

// GetUser is a helper function ontop of GetEntry
// which returns a pointer to a User
func (ut *UserTable) GetUser(userID string) *User {
	if ut.HasEntry(userID) {
		dte := ut.GetEntry(userID)
		return dte.(*User)
	}
	return nil
}

// SetUser is a helper function ontop of SetEntry
func (ut *UserTable) SetUser(usr *User) {
	ut.SetEntry(usr.UserID, usr)
}

// AddNewBlankUser creates a new user with a
// newly generated UUID.
func (ut *UserTable) AddNewBlankUser() *User {
	uuid, _ := uuid.NewV4()
	nwUser := User{}
	nwUser.UserID = uuid.String()
	ut.SetEntry(nwUser.UserID, &nwUser)
	return ut.GetUser(nwUser.UserID)
}

// CreateBlankEntry (IBlankTableEntryCreator interface)
func (ut *UserTable) CreateBlankEntry() DataTableEntry {
	return &User{}
}

//////////////////////////
