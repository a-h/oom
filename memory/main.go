package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

const increment = 1024 * 1024 * 256

func Handler() (Response, error) {
	fmt.Println("I'm about to use up a lot of RAM...")
	var space []byte

	for {
		space = append(space, make([]byte, increment)...)
		fmt.Printf("%dMB consumed\n", len(space)/1024/1024)
	}

	return Response{
		Message: "Unexpectedly large volume of RAM available...",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
