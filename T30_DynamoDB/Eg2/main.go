package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Person represents a row in the Person table.
type Person struct {
	PersonID  string
	FirstName string
	Surname   string
}

func main() {

	connConfig := DynamoDBConfigV1{
		acKey:     "",
		scKey:     "",
		awsRegion: "",
		endpoint:  "",
	}

	if connConfig.IsValid() == false {
		fmt.Println("ERROR: Database Conn Config was Invalid.")
		fmt.Println("Program Terminated.")
		return
	}

	var err error
	var dbConn *DynamoDBConnection

	dbConn, err = ConnectToDatabase(connConfig)
	if shouldContinue(err) {
		tableName := ""
		queryOutput, err := dbConn.GetTableRowByPartitionKey(tableName, "PersonID", "P2")
		if shouldContinue(err) {
			personList := []Person{}
			err = extractPersonList(queryOutput, &personList)
			if shouldContinue(err) {
				fmt.Println("Found this output:")
				fmt.Println(personList)
			}
		}
	}

	fmt.Println("Program Terminated.")
}

///////////////////////////////////////
// HELPER FUNCTIONS:
///////////////////////////////////////

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func shouldContinue(err error) bool {
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return false
	}
	return true
}

// TODO: Find a more generic way of doing this.
func extractPersonList(queryOutput *dynamodb.QueryOutput, retList *[]Person) error {

	if queryOutput == nil || retList == nil {
		return fmt.Errorf("extract-person-list failed")
	}

	err := dynamodbattribute.UnmarshalListOfMaps(queryOutput.Items, retList)
	if err != nil {
		return err
	}

	return nil
}

///////////////////////////////////////
