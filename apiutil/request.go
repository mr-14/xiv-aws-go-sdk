package apiutil

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Context defines service ctx
type Context struct {
	Params   map[string]string
	DynamoDB *dynamodb.DynamoDB
	// DynamoTX *dtx.TransactionManager
}

// HandlerSpec defines request handlers
type HandlerSpec struct {
	List    func(query map[string]string, ctx *Context) (events.APIGatewayProxyResponse, error)
	Add     func(body string, ctx *Context) (events.APIGatewayProxyResponse, error)
	Get     func(params map[string]string, ctx *Context) (events.APIGatewayProxyResponse, error)
	Edit    func(params map[string]string, body string, ctx *Context) (events.APIGatewayProxyResponse, error)
	Delete  func(params map[string]string, ctx *Context) (events.APIGatewayProxyResponse, error)
	DBEvent func(ctx context.Context, e events.DynamoDBEvent)
}

// Dispatch dispatches request to matching handler
func Dispatch(req events.APIGatewayProxyRequest, ctx *Context, handler HandlerSpec) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		if req.PathParameters != nil {
			return handler.Get(req.PathParameters, ctx)
		}
		return handler.List(req.QueryStringParameters, ctx)
	case "POST":
		return handler.Add(req.Body, ctx)
	case "PUT":
		return handler.Edit(req.PathParameters, req.Body, ctx)
	case "DELETE":
		return handler.Delete(req.PathParameters, ctx)
	}

	return NewMessageError(400, "error.request.invalid")
}
