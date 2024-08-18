package database

import (
	"handler/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


const TABLE_NAME = "Users"

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: db,
	}
}

// Does the user exist in the database?
// How do i insert a new record into DynamoDb

func (u DynamoDBClient) UserExists(username string) (bool, error) {
	// Do the check here
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return true, err
	}
	if result.Item == nil {
		return false, nil
	}
	return true, nil
}


func (u DynamoDBClient) InsertUser(user types.RegisterUser) error {
	// assemble the item then insert it
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}
	_, err := u.databaseStore.PutItem(item)

	if err != nil {
		return err
	}
	return nil
}