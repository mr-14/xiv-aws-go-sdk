package apiutil

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// NewErrorResponse creates error response
func NewErrorResponse(err error) events.APIGatewayProxyResponse {
	log.Printf("Error: %s", err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       `{"message": "error.internal"}`,
	}
}
