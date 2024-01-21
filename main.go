package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleLambdaEvent)
}

func handleLambdaEvent(event Event) (Response, error) {
	return Response{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}
