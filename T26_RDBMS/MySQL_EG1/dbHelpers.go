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

// RunMod runs the modification statement on the DBConnection.
// Mod statements are anything which modify the database.
// EG. CREATE, INSERT and DELETE.
func (dbConn *DBConnection) RunMod(mod string) {
	stmt, err := dbConn.RawConn.Prepare(mod)
	printOutErr(err)

	r, err := stmt.Exec()
	printOutErr(err)

	_, err = r.RowsAffected()
	printOutErr(err)
}

// RunQuery runs a query on the DBConnection.
// Queries do not modify the database.
func (dbConn *DBConnection) RunQuery(query string) (*sql.Rows, error) {
	rows, err := dbConn.RawConn.Query(query)
	return rows, err
}

// RunExec runs exec type SQL queries on the DBConnection.
func (dbConn *DBConnection) RunExec(query string) (sql.Result, error) {
	result, err := dbConn.RawConn.Exec(query)
	return result, err
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

// IsValidDBConnectionNoPtr similar to IsValidDBConnection but does not need a pointer.
func IsValidDBConnectionNoPtr(dbConn DBConnection) bool {
	return (dbConn.RawConn != nil)
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
	retConn.RawConn = db
	return retConn, retError
}

///////////////////////////////////////////
// OTHER HELPER FUNCTIONS:
///////////////////////////////////////////

func convertRowsToString(rows *sql.Rows) (string, error) {
	retStr := ""
	var retError error = nil
	if rows == nil {
		return "", fmt.Errorf("Rows was nil")
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	length := len(cols)

	var result [][]string
	var container []string
	var pointers []interface{}
	for rows.Next() {
		pointers = make([]interface{}, length)
		container = make([]string, length)
		for i := range pointers {
			pointers[i] = &container[i]
		}
		err = rows.Scan(pointers...)
		if err != nil {
			retStr = ""
			retError = err
			break
		}
		result = append(result, container)
	}

	if err != nil {
		retStr = ""
	}

	if err == nil {
		for idx, rowData := range result {
			for i := 0; i < len(rowData); i++ {
				retStr += rowData[i]
				if i < (len(rowData) - 1) {
					retStr += "\t"
				}
			}
			if idx < (len(result) - 1) {
				retStr += "\n"
			}
		}
	}

	return retStr, retError
}

///////////////////////////////////////////
