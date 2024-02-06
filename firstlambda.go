package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {

	result := "hello world"
	fmt.Println("request:", request.Body)

	return result, nil

}

func main() {
	lambda.Start(handler)
}
