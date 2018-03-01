package apiutil

import (
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mr-14/xiv-aws-go-sdk/errorutil"
)

// Container defines service container
type Container struct {
	Params   map[string]string
	DynamoDB *dynamodb.DynamoDB
	// DynamoTX *dtx.TransactionManager
}

// Route defines http path to handler mapping
type Route struct {
	Method  string
	Path    string
	Handler func(container *Container, req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse
}

// Dispatch dispatches request to matching handler
func Dispatch(container *Container, req events.APIGatewayProxyRequest, routes []*Route) (resp events.APIGatewayProxyResponse) {
	log.Printf("Request Received:\nPath: %s\nMethod: %s\nQuery: %+v\nBody: %s", req.Path, req.HTTPMethod, req.QueryStringParameters, req.Body)

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error: %s", err)
			resp = getErrorResponse(err)
		}
	}()

	for _, route := range routes {
		if route.Method != req.HTTPMethod {
			continue
		}

		if getPath(route.Path, req.PathParameters) == req.Path {
			resp = route.Handler(container, req)
			log.Printf("Response Sent:\nStatus: %d\nHeader: %+v\nBody: %s", resp.StatusCode, resp.Headers, resp.Body)
			return resp
		}
	}

	panic(errorutil.NewHTTPMessageError(400, "error.path.invalid"))
}

func getPath(path string, pathParams map[string]string) string {
	for key, val := range pathParams {
		path = strings.Replace(path, "{"+key+"}", val, 1)
	}

	return path
}

func getErrorResponse(err interface{}) events.APIGatewayProxyResponse {
	switch e := err.(type) {
	case *errorutil.FormError:
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       e.Error(),
		}
	case *errorutil.HTTPError:
		return events.APIGatewayProxyResponse{
			StatusCode: e.Status,
			Body:       e.Form.Error(),
		}
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"message": "error.internal"}`,
		}
	}
}
