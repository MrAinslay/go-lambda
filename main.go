package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	lambda.Start(handleLambdaEvent)
}

func handleLambdaEvent(event Event) (Response, error) {
	return Response{}, nil
}
