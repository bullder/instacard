package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleRequest(ctx context.Context) (MyResponse, error) {
	//return fmt.Sprintf("Hello %s!", name.Name ), nil
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", "Mike", 38)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
