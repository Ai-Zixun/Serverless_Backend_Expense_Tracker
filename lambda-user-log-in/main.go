package main

import (
	"fmt"
	"log"
	"time"

	encoder "expense-tracker/encoder"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Item - Database schema
type Item struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	CreationDate time.Time `json:"creation-time"`
}

// Request - Http request from the AWS API Gateway
type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response - Http response to the request from AWS API Gateway
type Response struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Ok      bool   `json:"ok"`
}

// DatabaseUserLogIn - Seeking matching item in the database
func DatabaseUserLogIn(username string, password string) bool {
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
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		fmt.Println("Error - Check User Exist ")
		log.Fatal(err)
		return false
	}

	if len(result.Item) == 0 {
		fmt.Println("User not exist in the database")
		return false
	}

	var item Item

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.Password == password {
		fmt.Println("Matching user item in the database")
		return true
	}
	fmt.Println("Missmatch password")
	return false

}

// Validate - Validate the username and password format
func Validate(username string, password string) (bool, string) {
	if len(username) > 15 || len(password) > 15 {
		return false, "Username or Password too long (Valid username and password should be between 5 - 15 letters)"
	}

	if len(username) < 5 || len(password) < 5 {
		return false, "Username or Password too short (Valid username and password should be between 5 - 15 letters)"
	}

	for _, char := range username {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false, "Invalid letter in username"
		}
	}

	for _, char := range password {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			return false, "Invalid letter in password"
		}
	}

	return true, ""
}

// Handler - Handler to process the AWS API Gateway request to create a user
func Handler(request Request) (Response, error) {

	if ok, msg := Validate(request.Username, request.Password); !ok {
		return Response{
			Message: msg,
			Key:     "null",
			Ok:      false,
		}, nil
	}

	key, _ := encoder.EncodeJWT(request.Username)

	if DatabaseUserLogIn(request.Username, request.Password) {
		return Response{
			Message: fmt.Sprintf("Log-in success"),
			Key:     key,
			Ok:      true,
		}, nil
	}

	return Response{
		Message: fmt.Sprintf("Log-in fail"),
		Key:     "null",
		Ok:      false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
