package main

import (
	"fmt"
	"handler/app"
	// "handler/app"

	"github.com/aws/aws-lambda-go/lambda"
)


type MyEvent struct {
	Username string `json:"username"`  // the json means that the field is marshalled/unmarshalled as JSON. marshalling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be moved between different parts of a computer program or from one program to another. unmarshalling is the reverse process, where a data stream created by serializing data is converted back into an object in memory.
}


func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is required")
	}
	return fmt.Sprintf("Hello %s!", event.Username), nil
	
}

func main() {

	app := app.NewApp()
	lambda.Start(app.ApiHandler.RegisterUserHandler)
	// lambda.Start(HandleRequest)

}

