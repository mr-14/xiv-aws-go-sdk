package apiutil

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mr-14/xiv-aws-go-sdk/errorutil"
)

// NewErrorResponse creates error response
func NewErrorResponse(err error) events.APIGatewayProxyResponse {
	log.Printf("Error: %s", err.Error())

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
