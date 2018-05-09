////////////////////////////////////////////////////
// DYNAMO DB UTILS:
////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	cErrorStringDBConnNil = "database connection was nil"
)

///////////////////////////////
// DynamoDBConfigV1
///////////////////////////////

// DynamoDBConfigV1 contains simplified settings needed to establish a connection with a DynamoDB database.
type DynamoDBConfigV1 struct {
	acKey     string
	scKey     string
	awsRegion string
	endpoint  string
}

// IsValid returns true if all config parameters are filled up correctly.
func (c *DynamoDBConfigV1) IsValid() bool {

	if c == nil {
		return false
	}

	return ((c.acKey != "") && (c.scKey != "") && (c.awsRegion != "") && (c.endpoint != ""))
}

///////////////////////////////
// DynamoDBConnection
///////////////////////////////

// DynamoDBConnection acts as a wrapper for dynamodb.DynamoDB (i.e. an active DynamoDB connection).
type DynamoDBConnection struct {
	AWSConn *dynamodb.DynamoDB
}

// IsValid returns true if the dbConn is valid and connected.
func (dbConn *DynamoDBConnection) IsValid() bool {
	if dbConn != nil {
		if dbConn.AWSConn != nil {
			return true
		}
	}
	return false
}

// ConnectToDatabase takes a connect config and attempts to establish a connection to a DynamoDB database.
// OnSuccess - Returns a pointer to a connected DynamoDB client and nil as error.
// OnFail - Returns nil and err info.
func ConnectToDatabase(config DynamoDBConfigV1) (*DynamoDBConnection, error) {

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
	return &DynamoDBConnection{AWSConn: dbConn}, err
}

// IsDBConnValid returns true if the passed in connection is valid and connected.
func IsDBConnValid(dbConn *DynamoDBConnection) bool {
	return ((dbConn != nil) && (dbConn.IsValid()))
}

// GetTableDescription gets the table description of a particular table in the connected DynamoDB database.
// OnSuccess - Returns a pointer to a TableDescription and nil as error.
// OnFail - Returns nil and error info.
func (dbConn *DynamoDBConnection) GetTableDescription(tableName string) (*dynamodb.TableDescription, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
	}

	request := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	result, err := dbConn.AWSConn.DescribeTable(request)
	var tableDesc *dynamodb.TableDescription = nil
	if err == nil {
		tableDesc = result.Table
	}
	return tableDesc, err
}

// GetAllRowsFromTable gets all data for all rows found within a DynamoDB database table.
func (dbConn *DynamoDBConnection) GetAllRowsFromTable(tableName string) (*dynamodb.ScanOutput, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dbConn.AWSConn.Scan(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTableRowByPartitionKey searches for a DB table item with the provided pk, in the table with the provided name.
func (dbConn *DynamoDBConnection) GetTableRowByPartitionKey(tableName string, pkColumnName, pkValue string) (*dynamodb.QueryOutput, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
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

	resp, err := dbConn.AWSConn.Query(queryInput)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetTableRowByGSI searches for a DB table item with the provided Global Secondary Index (GSI).
// The version of the function automatically sets maxRowsReturned to 1.
func (dbConn *DynamoDBConnection) GetTableRowByGSI(tableName string,
	gsiName string, gsiValue string, columnName string) (*dynamodb.QueryOutput, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
	}
	return dbConn.GetTableRowsByGSI(tableName, gsiName, gsiValue, columnName, 1)
}

// GetTableRowsByGSI searches for a DB table item with the provided Global Secondary Index (GSI).
// maxRowsReturned can be specified to set a limit on the number of rows returned.
func (dbConn *DynamoDBConnection) GetTableRowsByGSI(tableName string,
	gsiName string, gsiValue string, columnName string, maxRowsReturned int) (*dynamodb.QueryOutput, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
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

	resp, err := dbConn.AWSConn.Query(queryInput)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RunQueryOnTable executes a query on a DynamoDB Table and fills up the result structure with any returned rows.
func (dbConn *DynamoDBConnection) RunQueryOnTable(tableName, query *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {

	if IsDBConnValid(dbConn) == false {
		return nil, fmt.Errorf(cErrorStringDBConnNil)
	}

	resp, err := dbConn.AWSConn.Query(query)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

////////////////////////////////////////////////////
