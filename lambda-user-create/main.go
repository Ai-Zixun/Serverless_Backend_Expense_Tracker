package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserName string `json:"user-name"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message`
	Key     string `json:"key"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Hello %s, Your Password is: %s", request.UserName, request.Password),
		Key:     "null",
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
