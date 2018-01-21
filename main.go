package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() {
	log.Print("Hello Lambda, I'm here to rock!")
}

func main() {
	lambda.Start(HandleRequest)
}
