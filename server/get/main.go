package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func HandleRequest(ctx context.Context) (Response, error) {
	msg := fmt.Sprintf("%s is %d years old!", "Mike", 38)
	fmt.Println(msg)

	return Response{
			StatusCode: 1,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       msg,
		},
		nil
}

func main() {
	lambda.Start(HandleRequest)
}
