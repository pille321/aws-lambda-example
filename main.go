package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

//MyEvent input
type MyEvent struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

//MyResponse output
type MyResponse struct {
	Message string `json:"Answer:"`
}

//HandleLambdaEvent handles MyEvent
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() { lambda.Start(HandleLambdaEvent) }
