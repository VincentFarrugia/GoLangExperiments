package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// SQLTableEntry represents a base type for sql tables.
type SQLTableEntry interface {
	GetTableName() string
	GetColumnHeaders() []string
	GetValues() []interface{}
}

// ClearObjectData clears data from the provided implementor of the SQLTableEntry interface.
func ClearObjectData(t SQLTableEntry) {
	values := t.GetValues()
	for _, item := range values {
		switch v := item.(type) {
		case *string:
			*v = ""
		case *int:
			*v = 0
		case *float32:
			*v = 0.0
		case *float64:
			*v = 0.0
		default:
		}
	}
}

// CreateSQLColumnsString generates a comma separated columns string.
func CreateSQLColumnsString(t SQLTableEntry, bAddTableName bool, bEncloseInBrackets bool) string {
	var buffer bytes.Buffer
	columnHeaders := t.GetColumnHeaders()
	numColumnHeaders := len(columnHeaders)
	tableName := t.GetTableName()
	if bEncloseInBrackets {
		buffer.WriteString("(")
	}
	for idx, item := range columnHeaders {

		if bAddTableName {
			buffer.WriteString(fmt.Sprintf("%s.%s", tableName, item))
		} else {
			buffer.WriteString(item)
		}

		if idx < numColumnHeaders {
			buffer.WriteString(",")
		}
	}
	if bEncloseInBrackets {
		buffer.WriteString(")")
	}
	return buffer.String()
}

// CreateSQLValuesString generates a comma separated values string.
func CreateSQLValuesString(t SQLTableEntry) string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	sqlValuesArr := ConvertValuesToStringArr(t)
	numValues := len(sqlValuesArr)
	for idx, item := range sqlValuesArr {
		buffer.WriteString(item)
		if idx < numValues {
			buffer.WriteRune(',')
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

// CreateCombinedColumnValueSQLString generates a column=value comma separated string.
func CreateCombinedColumnValueSQLString(t SQLTableEntry, bAddTableName bool) string {

	var buffer bytes.Buffer
	tableName := t.GetTableName()
	columnHeaders := t.GetColumnHeaders()
	sqlValues := ConvertValuesToStringArr(t)
	numColumnHeaders := len(columnHeaders)

	for idx, item := range columnHeaders {

		if bAddTableName {
			buffer.WriteString(tableName)
			buffer.WriteString(".")
		}
		buffer.WriteString(item)
		buffer.WriteString("=")
		buffer.WriteString(sqlValues[idx])

		if idx < numColumnHeaders {
			buffer.WriteString(",")
		}
	}

	return buffer.String()
}

// ConvertValuesToStringArr creates a sql string representation for all values for a table entry.
func ConvertValuesToStringArr(t SQLTableEntry) []string {
	retStringArr := []string{}
	vals := t.GetValues()
	for _, item := range vals {
		switch v := item.(type) {
		case *string:
			retStringArr = append(retStringArr, fmt.Sprintf(`"%s"`, *v))
		case *int:
			retStringArr = append(retStringArr, strconv.Itoa(*v))
		case *float32:
			retStringArr = append(retStringArr, fmt.Sprintf("%f", *v))
		case *float64:
			retStringArr = append(retStringArr, fmt.Sprintf("%f", *v))
		default:
			// Skip.
		}
	}
	return retStringArr
}

//////////

// UpdateTableRow updates a row in the related table.
func UpdateTableRow(dbConn *DBConnection, rowData SQLTableEntry) error {

	if dbConn == nil {
		return fmt.Errorf("DBConn was nil")
	}

	if rowData == nil {
		return fmt.Errorf("RowData was nil")
	}

	sqlQueryStr := fmt.Sprintf(`
	INSERT INTO %s %s
	VALUES %s
	ON DUPLICATE KEY UPDATE
	%s;`,
		cProfileTableName,
		CreateSQLColumnsString(rowData, false, true),
		CreateSQLValuesString(rowData),
		CreateCombinedColumnValueSQLString(rowData, false))
	//fmt.Printf("SQL QUERY IS: '%s'\n", sqlQueryStr)
	err := dbConn.RunMod(sqlQueryStr)
	return err
}

//////////
