package main

import (
	"fmt"
	"reflect"
)

// Person represents a row in the Person table.
type Person struct {
	PersonID  string
	FirstName string
	LastName  string
}

func main() {

	connConfig := DynamoDBConfigV1{
		acKey:     "",
		scKey:     "",
		awsRegion: "",
		endpoint:  "",
	}

	if connConfig.IsValid() == false {
		fmt.Println("Database Conn Config was Invalid. Exiting...")
		return
	}

	dbConn, err := ConnectToDatabase(connConfig)
	panicIfErr(err)

	obj := []Person{}
	retSlice, bOk := takeSliceArg(obj)
	if bOk == false {
		fmt.Println("slice conversion not ok")
	} else {
		tableName := ""
		err = GetTableRowByPartitionKey(dbConn, tableName, "PersonID", "P1", &retSlice)
		panicIfErr(err)

		if err == nil {

			obj = make([]Person, len(retSlice))
			for i, x := range retSlice {
				obj[i] = x.(Person)
			}

			fmt.Println("Returned this data:")
			fmt.Println(obj)
		}
	}
}

///////////////////////////////////////
// HELPER FUNCTIONS:
///////////////////////////////////////

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

///////////////////////////////////////
