package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	name := request.QueryStringParameters["name"]
	if name == "" {
		return Response{
			StatusCode: 400,
			Body:       "Missing 'name' query parameter",
		}, nil
	}

	responseData := map[string]interface{}{
		"message": "Hello, this is a JSON response from AWS Lambda!" + name,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: 200,
		Body:       string(jsonData),
	}, nil
}

func main() {
	lambda.Start(Handler)
}