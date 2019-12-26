package main

import (
	"fmt"
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

	DatabaseCreateUser(item)

	return Response{
		Message: fmt.Sprintf("Hello %s, Your Password is: %s", request.Username, request.Password),
		Key:     "null",
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
