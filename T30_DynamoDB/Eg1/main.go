package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var acKey = ""
var scKey = ""
var awsRegion = ""
var endpoint = ""

// Person represents an entry in the Person database table.
type Person struct {
	PersonID  string
	FirstName string
	LastName  string
}

func main() {

	if (acKey == "") || (scKey == "") || (awsRegion == "") || (endpoint == "") {
		fmt.Println("Detected empty conn settings. Exiting...")
		return
	}

	// Create DB Conn Config.
	cred := credentials.NewStaticCredentials(acKey, scKey, "")
	awsConfig := &aws.Config{
		Credentials: cred,
		Region:      &awsRegion,
		Endpoint:    &endpoint,
	}

	// Connect to the DB.
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		panic(err)
	}
	dbConn := dynamodb.New(sess)

	// Perform experiments.
	getTableDescription(dbConn, "Person")
	getAllRowsFromTable(dbConn, "Person")
	queryTestByID(dbConn)
	queryTestByGSIIndex(dbConn)
	queryTestMaxResultLimit(dbConn)
	queryTestTableNameSearch(dbConn)
}

func getTableDescription(dbConn *dynamodb.DynamoDB, tableName string) {

	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("Person"),
	}

	result, err := dbConn.DescribeTable(req)
	panicIfError(err)

	tableDesc := result.Table
	fmt.Println()
	fmt.Printf("Here is the description of the %s table\n", tableName)
	fmt.Println(tableDesc)
}

func getAllRowsFromTable(dbConn *dynamodb.DynamoDB, tableName string) {
	params := &dynamodb.ScanInput{
		TableName: aws.String("Person"),
	}
	result, err := dbConn.Scan(params)
	panicIfError(err)

	obj := []Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	panicIfError(err)

	fmt.Println()
	fmt.Printf("Here are all the rows from the %s table\n", tableName)
	for _, item := range obj {
		fmt.Println(item)
	}
}

func queryTestByID(dbConn *dynamodb.DynamoDB) {

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String("Person"),
		KeyConditions: map[string]*dynamodb.Condition{
			"PersonID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String("0"),
					},
				},
			},
		},
	}

	resp, err := dbConn.Query(queryInput)
	logIfError(err)
	personObj := []Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &personObj)
	fmt.Println()
	fmt.Println("This is the result of queryTestOne\n")
	fmt.Println(personObj)
}

func queryTestByGSIIndex(dbConn *dynamodb.DynamoDB) {

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String("Person"),
		IndexName: aws.String("FirstName-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"FirstName": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("Cabal"),
					},
				},
			},
		},
	}

	resp, err := dbConn.Query(queryInput)
	logIfError(err)
	personObj := []Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &personObj)
	fmt.Println()
	fmt.Println("This is the result of queryTestTwo\n")
	fmt.Println(personObj)
}

func queryTestMaxResultLimit(dbConn *dynamodb.DynamoDB) {

	queryInput := &dynamodb.QueryInput{
		Limit:     aws.Int64(1),
		TableName: aws.String("Person"),
		IndexName: aws.String("FirstName-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"FirstName": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("Cabal"),
					},
				},
			},
		},
	}

	resp, err := dbConn.Query(queryInput)
	logIfError(err)
	personObj := []Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &personObj)
	fmt.Println()
	fmt.Println("This is the result of queryTestThree\n")
	fmt.Println(personObj)
}

func queryTestTableNameSearch(dbConn *dynamodb.DynamoDB) {
	obj := dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("Person"),
		Limit: aws.Int64(1),
	}

	pageNum := 0
	err := dbConn.ListTablesPages(&obj,
		func(page *dynamodb.ListTablesOutput, lastPage bool) bool {
			fmt.Println("I found myself", page)
			pageNum++
			fmt.Println("I found this too", page)
			return pageNum <= 5
		})

	logIfError(err)
}

func logIfError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
