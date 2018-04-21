package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

//////////////////////////
// DATA TABLE ENTRY
//////////////////////////

// DataTableEntry represents an entry in a DataTable.
// This can be seen as one "row".
type DataTableEntry interface {
	InitBlank()
	GetPrimaryKey() string
	SetFromCSVLine(csvLine []string)
	ToCSVLine() string
}

//////////////////////////
// DATA TABLE
//////////////////////////

// DataTable represents a table containing metadata and entry rows.
type DataTable struct {
	Name          string
	ColumnHeaders []string
	Rows          map[string]DataTableEntry
	IsDirty       bool
}

// InitBlank creates a blank table with a given name.
func (dt *DataTable) InitBlank(tableName string) {
	dt.Name = tableName
	dt.ColumnHeaders = make([]string, 0)
	dt.Rows = make(map[string]DataTableEntry)
	dt.IsDirty = false
}

// InitFromCSVFile sets the datatable state via CSV File.
func (dt *DataTable) InitFromCSVFile(tableName string, csvFileFullPath string, entryCreator IBlankTableEntryCreator) {

	dt.InitBlank(tableName)

	if _, err := os.Stat(csvFileFullPath); os.IsNotExist(err) {
		// csv file does not exist.
		return
	}

	file, err := os.Open(csvFileFullPath)
	if err != nil {
		// Error occurred.
		// TODO: return error.
	}

	r := csv.NewReader(file)
	lineIdx := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// Some error occurred.
			// TODO: Return custom error type.
		} else {
			if lineIdx == 0 {
				// Read in header columns.
				dt.ColumnHeaders = append(dt.ColumnHeaders, record...)
			} else {
				// Read in row entry.
				dte := entryCreator.CreateBlankEntry()
				if dte != nil {
					dte.SetFromCSVLine(record)
					dt.SetEntry(dte.GetPrimaryKey(), dte)
				}
			}
		}
		lineIdx++
	}
}

// SaveToCSVFile saves the table's data to a csv file on disk.
func (dt *DataTable) SaveToCSVFile(csvFileFullPath string) {
	file, err := os.Create(csvFileFullPath)
	defer file.Close()
	if err != nil {
		// TODO return error.
		return
	}
	// Write Header Columns.
	columnHeaderStr := ""
	for idx, item := range dt.ColumnHeaders {
		columnHeaderStr += item
		if idx < (len(dt.ColumnHeaders) - 1) {
			columnHeaderStr += ","
		}
	}
	fmt.Fprintln(file, columnHeaderStr)
	// Write row entries.
	rowNumIdx := 0
	for _, rowEntry := range dt.Rows {
		entryAsCSV := rowEntry.ToCSVLine()
		if entryAsCSV != "" {
			if rowNumIdx < (len(dt.Rows) - 1) {
				entryAsCSV += ","
			}
			fmt.Fprintln(file, entryAsCSV)
		}
		rowNumIdx++
	}

	file.Close()

	dt.IsDirty = false
}

// SaveToCSVFileIfDirty only saves the datatable if it has been
// marked as modified (dirty).
func (dt *DataTable) SaveToCSVFileIfDirty(csvFileFullPath string) {
	if dt.IsDirty {
		dt.SaveToCSVFile(csvFileFullPath)
		dt.IsDirty = false
	}
}

// HasEntry returns true if the table contains
// an entry with the provided primary key.
func (dt *DataTable) HasEntry(pk string) bool {
	if _, bOk := dt.Rows[pk]; bOk {
		return true
	}
	return false
}

// GetEntry returns an entry with the matching primary key
// or nil if no such entry exits.
func (dt *DataTable) GetEntry(pk string) DataTableEntry {
	if entry, bOk := dt.Rows[pk]; bOk {
		return entry
	}
	return nil
}

// SetEntry inserts or overwrites an entry into the table
// and assigns it with the given primary key.
func (dt *DataTable) SetEntry(pk string, entry DataTableEntry) {
	if pk != "" {
		dt.Rows[pk] = entry
		dt.SetDirty()
	}
}

// RemoveEntry deletes an entry with the given primary key
// if it exists within the table.
func (dt *DataTable) RemoveEntry(pk string) {
	if dt.HasEntry(pk) {
		delete(dt.Rows, pk)
		dt.SetDirty()
	}
}

// SetDirty marks a datatable as modified (dirty)
// this is useful for external users to determine
// if they would like to update other instances of the
// same table. Eg. save the updated table to disk.
func (dt *DataTable) SetDirty() {
	dt.IsDirty = true
}

// CreateBlankEntry (IBlankTableEntryCreator interface)
// Should be implemented in child structs of DataTable.
func (dt *DataTable) CreateBlankEntry() DataTableEntry {
	return nil
}

////////////////////////////
// IBlankTableEntryCreator
////////////////////////////

// IBlankTableEntryCreator interface for creating a concrete type of a table entry.
type IBlankTableEntryCreator interface {
	CreateBlankEntry() DataTableEntry
}

////////////////////////////
