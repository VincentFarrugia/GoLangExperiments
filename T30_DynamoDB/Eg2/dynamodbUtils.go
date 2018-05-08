////////////////////////////////////////////////////
// DYNAMO DB UTILS:
////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	cErrorStringDBConnNil = "database connection was nil"
)

// DynamoDBConfigV1 contains simplified settings needed to establish a connection with a DynamoDB database.
type DynamoDBConfigV1 struct {
	acKey     string
	scKey     string
	awsRegion string
	endpoint  string
}

func (c *DynamoDBConfigV1) IsValid() bool {

	if c == nil {
		return false
	}

	return ((c.acKey != "") && (c.scKey != "") && (c.awsRegion != "") && (c.endpoint != ""))
}

// ConnectToDatabase takes a connect config and attempts to establish a connection to a DynamoDB database.
// OnSuccess - Returns a pointer to a connected DynamoDB client and nil as error.
// OnFail - Returns nil and err info.
func ConnectToDatabase(config DynamoDBConfigV1) (*dynamodb.DynamoDB, error) {

	if config.IsValid() == false {
		return nil, fmt.Errorf("database conn config is invalid")
	}

	// Create AWS credentials from simple config.
	cred := credentials.NewStaticCredentials(config.acKey, config.scKey, "")
	awsConfig := &aws.Config{
		Credentials: cred,
		Region:      &config.awsRegion,
		Endpoint:    &config.endpoint,
	}

	// Connect to the DB.
	sess, err := session.NewSession(awsConfig)
	var dbConn *dynamodb.DynamoDB = nil
	if err == nil {
		dbConn = dynamodb.New(sess)
	}
	return dbConn, err
}

// GetTableDescription gets the table description of a particular table in the connected DynamoDB database.
// OnSuccess - Returns a pointer to a TableDescription and nil as error.
// OnFail - Returns nil and error info.
func GetTableDescription(dbConn *dynamodb.DynamoDB, tableName string) (*dynamodb.TableDescription, error) {

	if dbConn == nil {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
	}

	request := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	result, err := dbConn.DescribeTable(request)
	var tableDesc *dynamodb.TableDescription = nil
	if err == nil {
		tableDesc = result.Table
	}
	return tableDesc, err
}

// GetAllRowsFromTable gets all data for all rows found within a DynamoDB database table.
func GetAllRowsFromTable(dbConn *dynamodb.DynamoDB, tableName string, retItems *[]interface{}) error {

	if dbConn == nil {
		return fmt.Errorf(cErrorStringDBConnNil)
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dbConn.Scan(params)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, retItems)
	if err != nil {
		return err
	}

	return nil
}

// GetTableRowByPartitionKey searches for a DB table item with the provided pk, in the table with the provided name.
func GetTableRowByPartitionKey(dbConn *dynamodb.DynamoDB, tableName string, pkColumnName, pkValue string, retItem *[]interface{}) error {

	if dbConn == nil {
		return fmt.Errorf(cErrorStringDBConnNil)
	}

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			pkColumnName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(pkValue),
					},
				},
			},
		},
	}

	resp, err := dbConn.Query(queryInput)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, retItem)
	return err
}

// GetTableRowByGSI searches for a DB table item with the provided Global Secondary Index (GSI).
// The version of the function automatically sets maxRowsReturned to 1.
func GetTableRowByGSI(dbConn *dynamodb.DynamoDB, tableName string,
	gsiName string, gsiValue string, columnName string,
	retItem *[]interface{}) error {
	return GetTableRowsByGSI(dbConn, tableName, gsiName, gsiValue, columnName, 1, retItem)
}

// GetTableRowsByGSI searches for a DB table item with the provided Global Secondary Index (GSI).
// maxRowsReturned can be specified to set a limit on the number of rows returned.
func GetTableRowsByGSI(dbConn *dynamodb.DynamoDB, tableName string,
	gsiName string, gsiValue string, columnName string, maxRowsReturned int,
	retItems *[]interface{}) error {

	if dbConn == nil {
		return fmt.Errorf(cErrorStringDBConnNil)
	}

	queryInput := &dynamodb.QueryInput{
		Limit:     aws.Int64(int64(maxRowsReturned)),
		TableName: aws.String(tableName),
		IndexName: aws.String(gsiName),
		KeyConditions: map[string]*dynamodb.Condition{
			columnName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(gsiValue),
					},
				},
			},
		},
	}

	resp, err := dbConn.Query(queryInput)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, retItems)
	return err
}

// RunQueryOnTable executes a query on a DynamoDB Table and fills up the result structure with any returned rows.
func RunQueryOnTable(dbConn *dynamodb.DynamoDB, tableName, query *dynamodb.QueryInput, retItems *[]interface{}) error {

	if dbConn == nil {
		return fmt.Errorf(cErrorStringDBConnNil)
	}

	resp, err := dbConn.Query(query)
	if err != nil {
		return err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, retItems)
	return err
}

////////////////////////////////////////////////////
