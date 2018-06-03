package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
)

// SQLTableEntry represents a base type for sql tables.
type SQLTableEntry interface {
	GetTableName() string
	GetOwnerDatabaseName() string
	GetPrimaryKeyName() string
	GetColumnHeaders() []string
	GetValues() []interface{}
	GenerateNewItem() SQLTableEntry
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
		case *bool:
			*v = false
		default:
		}
	}
}

// CreateSQLColumnsString generates a comma separated columns string.
func CreateSQLColumnsString(t SQLTableEntry, bAddTableName bool, bAddDatabaseName, bEncloseInBrackets bool) string {
	var buffer bytes.Buffer
	columnHeaders := t.GetColumnHeaders()
	numColumnHeaders := len(columnHeaders)
	tableName := t.GetTableName()
	dbName := t.GetOwnerDatabaseName()
	if bEncloseInBrackets {
		buffer.WriteString("(")
	}
	for idx, item := range columnHeaders {

		if bAddDatabaseName && bAddTableName {
			buffer.WriteString(fmt.Sprintf("%s.%s.%s", dbName, tableName, item))
		} else if bAddTableName {
			buffer.WriteString(fmt.Sprintf("%s.%s", tableName, item))
		} else {
			buffer.WriteString(item)
		}

		if idx < (numColumnHeaders - 1) {
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
		if idx < (numValues - 1) {
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

		if idx < (numColumnHeaders - 1) {
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
		case *bool:
			retStringArr = append(retStringArr, fmt.Sprintf("%t", *v))
		default:
			// Skip.
		}
	}
	return retStringArr
}

//////////

// GetTableRowsByPK attempts to get table rows using the provided primary keys.
func GetTableRowsByPK(dbConn *DBConnection, primaryKeyList []string, dummyItem SQLTableEntry, oItems *[]SQLTableEntry) (*[]SQLTableEntry, error) {

	if dbConn == nil {
		return oItems, fmt.Errorf("DBConn was nil")
	}

	if oItems == nil {
		return nil, fmt.Errorf("oItems was nil")
	}

	dbTableCombo := dummyItem.GetOwnerDatabaseName() + "." + dummyItem.GetTableName()

	sqlQueryStr := fmt.Sprintf(`
	SELECT *
	FROM %s 
	`,
		dbTableCombo)

	var buffer bytes.Buffer
	buffer.WriteString("WHERE ")
	numIDs := len(primaryKeyList)
	for idx, item := range primaryKeyList {
		buffer.WriteString(fmt.Sprintf(`%s.%s="`, dummyItem.GetTableName(), dummyItem.GetPrimaryKeyName()))
		buffer.WriteString(item)
		buffer.WriteString(`"`)
		if idx < (numIDs - 1) {
			buffer.WriteString("||")
		}
	}
	buffer.WriteString(";")
	sqlQueryStr += buffer.String()

	sqlRows, err := dbConn.RunQuery(sqlQueryStr)
	if err != nil {
		return oItems, err
	}

	retItems, err := ConvertSQLRowsToModelObjectSlice(sqlRows, dummyItem, oItems)
	return retItems, err
}

// UpdateTableRow updates a row in the related table.
func UpdateTableRow(dbConn *DBConnection, rowData SQLTableEntry) error {
	return UpdateTableRowV2(dbConn, rowData, "")
}

// UpdateTableRowV2 updates a row in the related table. This variant includes a custom "onDuplicateUpdateStr" string.
func UpdateTableRowV2(dbConn *DBConnection, rowData SQLTableEntry, onDuplicateUpdateStr string) error {

	if dbConn == nil {
		return fmt.Errorf("DBConn was nil")
	}

	if rowData == nil {
		return fmt.Errorf("RowData was nil")
	}

	dupUpdateStr := onDuplicateUpdateStr
	if onDuplicateUpdateStr == "" {
		dupUpdateStr = CreateCombinedColumnValueSQLString(rowData, false)
	}

	sqlQueryStr := fmt.Sprintf(`
	INSERT INTO %s %s
	VALUES %s
	ON DUPLICATE KEY UPDATE 
	%s;`,
		rowData.GetTableName(),
		CreateSQLColumnsString(rowData, false, false, true),
		CreateSQLValuesString(rowData),
		dupUpdateStr)
	//fmt.Println()
	//fmt.Printf("SQL QUERY IS: '%s'\n", sqlQueryStr)
	//fmt.Println()
	err := dbConn.RunMod(sqlQueryStr)
	return err
}

// BulkUpdateTableRows performs a batch update on multiple rows for a particular table.
func BulkUpdateTableRows(dbConn *DBConnection, rowDataList []SQLTableEntry) error {

	if dbConn == nil {
		return fmt.Errorf("DBConn was nil")
	}

	if rowDataList == nil {
		return fmt.Errorf("RowDataList was nil")
	}

	numRows := len(rowDataList)
	if numRows <= 0 {
		// RowData was empty.
		return nil
	}

	var strBuffer bytes.Buffer
	columnHeaderStr := CreateSQLColumnsString(rowDataList[0], false, false, true)
	strBuffer.WriteString(fmt.Sprintf("INSERT INTO %s %s VALUES ", rowDataList[0].GetTableName(), columnHeaderStr))

	for idx, item := range rowDataList {
		strBuffer.WriteString(CreateSQLValuesString(item))

		if idx < (numRows - 1) {
			strBuffer.WriteString(",")
		}
		strBuffer.WriteString("\n")
	}
	strBuffer.WriteString(" ON DUPLICATE KEY UPDATE ")
	columnHeaders := rowDataList[0].GetColumnHeaders()
	numColumns := len(columnHeaders)
	for idx, item := range columnHeaders {
		strBuffer.WriteString(fmt.Sprintf("%s=VALUES(%s)", item, item))

		if idx < (numColumns - 1) {
			strBuffer.WriteString(",")
		}
		strBuffer.WriteString("\n")
	}
	strBuffer.WriteString(";")

	sqlQueryStr := strBuffer.String()
	fmt.Println("QueryString is:" + sqlQueryStr)
	err := dbConn.RunMod(sqlQueryStr)
	return err
}

// DoesRowExistWithStringPK returns true if the row with the provided PK exists in the specificed Database.Table.
func DoesRowExistWithStringPK(dbConn *DBConnection, databaseName string, tableName string, pkName string, pkValue string) (bool, error) {

	if dbConn == nil {
		return true, fmt.Errorf("DBConn was nil")
	}

	dbTableCombo := databaseName + "." + tableName

	sqlQueryStr := fmt.Sprintf(`
	SELECT EXISTS( SELECT 1 FROM %s WHERE %s.%s="%s")`,
		dbTableCombo,
		dbTableCombo,
		pkName,
		pkValue)

	res, err := dbConn.RunExec(sqlQueryStr)
	if err != nil {
		return true, err
	}

	numRowsFound, err := res.RowsAffected()
	if err != nil {
		return true, err
	}

	return (numRowsFound > 0), nil
}

// GetTableRowCount attempts to get the size of the table in terms of number of rows.
func GetTableRowCount(dbConn *DBConnection, databaseName string, tableName string) (int, error) {

	if dbConn == nil {
		return 0, fmt.Errorf("DBConn was nil")
	}

	dbTableCombo := databaseName + "." + tableName

	sqlQueryStr := fmt.Sprintf(`
		SELECT COUNT(1) FROM %s`,
		dbTableCombo)

	res, err := dbConn.RunQuery(sqlQueryStr)
	if err != nil {
		return 0, err
	}

	retCount := 0
	if res.Next() {
		err = res.Scan(&retCount)
		if err != nil {
			return 0, err
		}
	}

	return retCount, nil
}

// GetAllTableData fills up the provided slice with all the rows for a particular table.
// Be very careful when using this as it will consume a lot of memory if the table is large.
// TODO: Make an alternate version of this which uses PAGING.
func GetAllTableData(dbConn *DBConnection, databaseName string, tableName string, dummyModelItem SQLTableEntry, oRetRows *[]SQLTableEntry) (*[]SQLTableEntry, error) {

	if dbConn == nil {
		return oRetRows, fmt.Errorf("DBConn was nil")
	}

	if oRetRows == nil {
		return oRetRows, fmt.Errorf("oRetRows was nil")
	}

	dbTableCombo := databaseName + "." + tableName

	sqlQueryStr := fmt.Sprintf(`
		SELECT * FROM %s`,
		dbTableCombo)

	sqlRows, err := dbConn.RunQuery(sqlQueryStr)
	if err != nil {
		return oRetRows, err
	}

	rowRes, err := ConvertSQLRowsToModelObjectSlice(sqlRows, dummyModelItem, oRetRows)
	return rowRes, err
}

// ConvertSQLRowsToModelObjectSlice attempts to create a model object for each row found in the sql rows result.
func ConvertSQLRowsToModelObjectSlice(sqlRows *sql.Rows, dummyModelItem SQLTableEntry, oRetDataSlice *[]SQLTableEntry) (*[]SQLTableEntry, error) {

	if oRetDataSlice == nil {
		return oRetDataSlice, fmt.Errorf("oRetDataSlice was nil")
	}

	if sqlRows == nil {
		// Row results was empty.
		return nil, nil
	}

	tmpSlice := *oRetDataSlice

	for sqlRows.Next() {

		nxtModelItem := dummyModelItem.GenerateNewItem()
		values := nxtModelItem.GetValues()
		valueSlice := make([]interface{}, 0)
		valueSlice = append(valueSlice, values...)

		err := sqlRows.Scan(valueSlice...)
		if shouldBreak(err) {
			continue
		} else {
			tmpSlice = append(tmpSlice, nxtModelItem)
		}
	}

	oRetDataSlice = &tmpSlice

	return &tmpSlice, nil
}

//////////
