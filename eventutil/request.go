package eventutil

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Container defines service container
type Container struct {
	Params   map[string]string
	DynamoDB *dynamodb.DynamoDB
}

// HandlerSpec defines request handlers
type HandlerSpec struct {
	Add    func(container *Container, record *events.DynamoDBStreamRecord)
	Edit   func(container *Container, record *events.DynamoDBStreamRecord)
	Delete func(container *Container, record *events.DynamoDBStreamRecord)
}

// Dispatch dispatches request to matching handler
func Dispatch(e events.DynamoDBEvent, container *Container, handler HandlerSpec) {
	for _, record := range e.Records {
		log.Printf("Processing request data for event ID %s.\n", record.EventID)
		switch record.EventName {
		case "INSERT":
			handler.Add(container, &record.Change)
		case "MODIFY":
			handler.Edit(container, &record.Change)
		case "REMOVE":
			handler.Delete(container, &record.Change)
		default:
			log.Printf("Stream event not supported: %s", record.EventName)
		}
	}
}
