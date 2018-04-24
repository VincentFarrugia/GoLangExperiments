package main

import (
	"database/sql"
	"fmt"
)

///////////////////////////////////////////
// DBConnSettings
///////////////////////////////////////////

// DBConnSettings is a helper struct
// for storing database connection settings.
type DBConnSettings struct {
	DriverName    string
	DBEndpoint    string
	DBPort        uint
	DBTableName   string
	DBUser        string
	DBUserShadow  string
	CharSetHeader string
}

// GetDataSourceName creates a connection string from the conn settings.
func (dbcs *DBConnSettings) GetDataSourceName() string {

	if dbcs == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		dbcs.DBUser,
		dbcs.DBUserShadow,
		dbcs.DBEndpoint,
		dbcs.DBPort,
		dbcs.DBTableName,
		dbcs.CharSetHeader)
}

// IsValid returns true if the DB connection settings are valid.
// Currently we only check if all the parameters are filled.
func (dbcs *DBConnSettings) IsValid() bool {

	if dbcs == nil {
		return false
	}

	return ((dbcs.DriverName != "") && (dbcs.DBEndpoint != "") && (dbcs.DBTableName != "") && (dbcs.DBUser != "") && (dbcs.DBUserShadow != "") && (dbcs.CharSetHeader != ""))
}

///////////////////////////////////////////
// DBConnection
///////////////////////////////////////////

type DBConnection struct {
	RawConn *sql.DB
}

// Close is a helper function for closing a connection to a DB.
func (dbConn *DBConnection) Close() error {
	var err error = nil
	if IsValidDBConnection(dbConn) {
		err = dbConn.RawConn.Close()
		dbConn.RawConn = nil
	}
	return err
}

// Ping is a helper function for pinging a DB using an existing connection.
func (dbConn *DBConnection) Ping() error {
	if IsValidDBConnection(dbConn) {
		return dbConn.RawConn.Ping()
	}
	return nil
}

// Invalidate makes a DBConnection unusable by closing it and resetting it.
func (dbConn *DBConnection) Invalidate() error {
	if IsValidDBConnection(dbConn) {
		err := dbConn.RawConn.Close()
		if err == nil {
			dbConn.RawConn = nil
		}
		return err
	}
	return nil
}

// IsValidDBConnection returns true if a connection is not null.
func IsValidDBConnection(dbConn *DBConnection) bool {
	return ((dbConn != nil) && (dbConn.RawConn != nil))
}

///////////////////////////////////////////
// DBConnCreator
///////////////////////////////////////////

// DBConnCreator creates DBConnections.
type DBConnCreator struct {
}

// CreateConnection connects to a DB using provided settings and
// automatically performs a ping.
func (dbcc *DBConnCreator) CreateConnection(settings DBConnSettings) (DBConnection, error) {

	retConn := DBConnection{RawConn: nil}
	var retError error = nil

	db, err := sql.Open(settings.DriverName, settings.GetDataSourceName())
	if err != nil {
		retConn.Invalidate()
		retError = err
		return retConn, retError
	}
	err = db.Ping()
	if err != nil {
		retConn.Invalidate()
		retError = err
		return retConn, retError
	}
	return retConn, retError
}

///////////////////////////////////////////
