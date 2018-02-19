package apiutil

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Container defines service container
type Container struct {
	Params   map[string]string
	DynamoDB *dynamodb.DynamoDB
	// DynamoTX *dtx.TransactionManager
}

// HandlerSpec defines request handlers
type HandlerSpec struct {
	List   func(container *Container, query map[string]string) (events.APIGatewayProxyResponse, error)
	Add    func(container *Container, body string) (events.APIGatewayProxyResponse, error)
	Get    func(container *Container, params map[string]string) (events.APIGatewayProxyResponse, error)
	Edit   func(container *Container, params map[string]string, body string) (events.APIGatewayProxyResponse, error)
	Delete func(container *Container, params map[string]string) (events.APIGatewayProxyResponse, error)
}

// Dispatch dispatches request to matching handler
func Dispatch(req events.APIGatewayProxyRequest, container *Container, handler HandlerSpec) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request Method: %s\nRequest Path: %s\nRequest Body: %s", req.HTTPMethod, req.Path, req.Body)

	switch req.HTTPMethod {
	case "GET":
		if req.PathParameters != nil {
			return handler.Get(container, req.PathParameters)
		}
		return handler.List(container, req.QueryStringParameters)
	case "POST":
		return handler.Add(container, req.Body)
	case "PUT":
		return handler.Edit(container, req.PathParameters, req.Body)
	case "DELETE":
		return handler.Delete(container, req.PathParameters)
	}

	return NewMessageError(400, "error.request.invalid")
}
