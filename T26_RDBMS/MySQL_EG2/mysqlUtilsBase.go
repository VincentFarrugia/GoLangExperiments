////////////////////////////////////////////////////
// MySQL DATABASE UTILS BASE:
////////////////////////////////////////////////////

package mysqlutils

import (
	"database/sql"
	"fmt"
	"time"
)

///////////////////////////////////////////
// DBConnSettings
///////////////////////////////////////////

// DBConnSettings is a helper struct
// for storing database connection settings.
type DBConnSettings struct {
	DriverName   string
	DBEndpoint   string
	DBPort       uint
	DBUser       string
	DBUserShadow string
	DBName       string
}

// GetDataSourceName creates a connection string from the conn settings.
func (dbcs *DBConnSettings) GetDataSourceName() string {

	if dbcs == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbcs.DBUser,
		dbcs.DBUserShadow,
		dbcs.DBEndpoint,
		dbcs.DBPort,
		dbcs.DBName)
}

// IsValid returns true if the DB connection settings are valid.
// Currently we only check if all the parameters are filled.
func (dbcs *DBConnSettings) IsValid() bool {

	if dbcs == nil {
		return false
	}

	return ((dbcs.DriverName != "") && (dbcs.DBEndpoint != "") && (dbcs.DBUser != "") && (dbcs.DBUserShadow != "") && (dbcs.DBName != ""))
}

///////////////////////////////////////////
// DBConnection
///////////////////////////////////////////

// DBConnection acts as a wrapper around a raw database connection.
type DBConnection struct {
	RawConn *sql.DB
}

// Close is a helper function for closing a connection to a DB.
func (dbConn *DBConnection) Close() error {
	var err error
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
func (dbConn *DBConnection) RunMod(mod string) error {
	stmt, err := dbConn.RawConn.Prepare(mod)
	if err != nil {
		return err
	}

	r, err := stmt.Exec()
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()
	return err
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
	var retError error

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
// QueryResult
///////////////////////////////////////////

// IQueryResult should be implemented by all structs used to parse a series of sql.Rows returned by SQL queries.
type IQueryResult interface {
	GetValues() []interface{}
	GenerateNewItem() IQueryResult
}

// ConvertRowsToQueryResultSlice transforms raw sql.Rows into a slice of IQueryResult.
func ConvertRowsToQueryResultSlice(rows *sql.Rows, dummyItem IQueryResult) ([]IQueryResult, error) {
	retSlice := []IQueryResult{}
	if rows == nil {
		return nil, fmt.Errorf("Rows was nil")
	}

	for rows.Next() {
		nxtItem := dummyItem.GenerateNewItem()
		valueSlice := make([]interface{}, 0)
		valueSlice = append(valueSlice, nxtItem.GetValues()...)

		err := rows.Scan(valueSlice...)
		if err != nil {
			continue
		} else {
			retSlice = append(retSlice, nxtItem)
		}
	}

	return retSlice, nil
}

///////////////////////////////////////////
// HELPER FUNCTIONS:
///////////////////////////////////////////

// ConvertRowsToString writes data from raw sql.Rows into a string and returns it.
func ConvertRowsToString(rows *sql.Rows) (string, error) {
	retStr := ""
	var retError error
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

// ConvertISO8601ToMySQLDatetime is a helper function to convert
// an ISO8601 DateTime format: Eg. "2018-05-15T09:52:21Z"
// into a MySQL Datetime type: Eg. "2018-05-15 09:52:21"
func ConvertISO8601ToMySQLDatetime(isoStr string) string {
	retStr := isoStr[0:10]
	retStr += " "
	retStr += isoStr[11:19]
	return retStr
}

// ConvertMySQLDatetimeToISO8601 is a helper function to convert
// a MySQL Datetime type: Eg. "2018-05-15 09:52:21"
// into an ISO8601 DateTime format: Eg. "2018-05-15T09:52:21Z"
func ConvertMySQLDatetimeToISO8601(mySQLDateTimeStr string) string {
	retStr := mySQLDateTimeStr[0:10]
	retStr += "T"
	retStr += mySQLDateTimeStr[11:]
	retStr += "Z"
	return retStr
}

// ExtractDateStringFromMySQLDatetime is a helper function to get
// the date portion of a MySQL Datetime type:
// Eg. FROM "2018-05-15 09:52:21" TO "2018-05-15"
func ExtractDateStringFromMySQLDatetime(mySQLDateTimeStr string) string {
	return mySQLDateTimeStr[0:10]
}

// ExtractDateStringFromISO8601 is a helper function to get
// the date portion of an ISO8601 datetime:
// EG. FROM "2018-05-15T09:52:21Z" TO "2018-05-15"
func ExtractDateStringFromISO8601(isoStr string) string {
	return isoStr[0:10]
}

// GetTimeUTCNowAsISO8601 returns the current time in UTC in the ISO8601 date-time format.
func GetTimeUTCNowAsISO8601() string {
	return time.Now().UTC().Format(time.RFC3339)
}

///////////////////////////////////////////
