package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation-time"`
}

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message`
	Key     string `json:"key"`
	Ok      bool   `json:"ok"`
}

func DatabaseUserExist(item Item) bool {
	// Load Credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB Client
	svc := dynamodb.New(sess)

	tableName := "Expense-Tracker-Users"
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(item.Username),
			},
		},
	})

	if err != nil {
		fmt.Println("Error - Check User Exist ")
		log.Fatal(err)
		return true
	}

	if len(result.Item) != 0 {
		fmt.Println(result.Item)
		fmt.Println("DynamoDB Username Found")
		return true
	}

	fmt.Println("DynamoDB Username Available")
	return false
}

func DatabaseCreateUser(item Item) {
	// Load Credentials
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// DynamoDB Client
	svc := dynamodb.New(sess)

	// Error handler
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tableName := "Expense-Tracker-Users"
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("DynamoDB User Create Complete")
}

func Handler(request Request) (Response, error) {

	item := Item{
		Username:     request.Username,
		Password:     request.Password,
		CreationDate: time.Now(),
	}

	// Return if the username is taken
	if DatabaseUserExist(item) {
		return Response{
			Message: fmt.Sprintf("The username %s is taken", request.Username),
			Key:     "null",
			Ok:      false,
		}, nil
	}

	DatabaseCreateUser(item)

	return Response{
		Message: fmt.Sprintf("Hello %s, Your user is created", request.Username),
		Key:     "null",
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
