package httputil

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Context defines service container
type Context struct {
	DynamoDB *dynamodb.DynamoDB
}

// HandlerSpec defines request handlers
type HandlerSpec struct {
	List   func(query map[string]string, container *Context) (events.APIGatewayProxyResponse, error)
	Add    func(body string, container *Context) (events.APIGatewayProxyResponse, error)
	Get    func(params map[string]string, container *Context) (events.APIGatewayProxyResponse, error)
	Edit   func(params map[string]string, body string, container *Context) (events.APIGatewayProxyResponse, error)
	Delete func(params map[string]string, container *Context) (events.APIGatewayProxyResponse, error)
}

// Dispatch dispatches request to matching handler
func Dispatch(req events.APIGatewayProxyRequest, container *Context, handler HandlerSpec) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		if req.PathParameters != nil {
			return handler.Get(req.PathParameters, container)
		}
		return handler.List(req.QueryStringParameters, container)
	case "POST":
		return handler.Add(req.Body, container)
	case "PUT":
		return handler.Edit(req.PathParameters, req.Body, container)
	case "DELETE":
		return handler.Delete(req.PathParameters, container)
	}

	return NewMessageError(400, "error.request.invalid")
}
