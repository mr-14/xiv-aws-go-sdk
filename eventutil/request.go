package eventutil

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Context defines service ctx
type Context struct {
	Params   map[string]string
	DynamoDB *dynamodb.DynamoDB
}

// HandlerSpec defines request handlers
type HandlerSpec struct {
	Add    func(ctx *Context, e events.DynamoDBEvent)
	Edit   func(ctx *Context, e events.DynamoDBEvent)
	Delete func(ctx *Context, e events.DynamoDBEvent)
}

// Dispatch dispatches request to matching handler
func Dispatch(e events.DynamoDBEvent, ctx *Context, handler HandlerSpec) {
	// switch req.HTTPMethod {
	// case "GET":
	// 	if req.PathParameters != nil {
	// 		handler.Get(req.PathParameters, ctx)
	// 	}
	// 	handler.List(req.QueryStringParameters, ctx)
	// case "POST":
	// 	handler.Add(req.Body, ctx)
	// case "PUT":
	// 	handler.Edit(req.PathParameters, req.Body, ctx)
	// case "DELETE":
	// 	handler.Delete(req.PathParameters, ctx)
	// }
}
