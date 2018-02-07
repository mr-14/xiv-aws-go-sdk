package httputil

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorResponse struct {
	Message string        `json:"message,omitempty"`
	Fields  []*ErrorField `json:"fields,omitempty"`
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewMessageError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       message,
	}, nil
}

func NewFieldError(statusCode int, fields []*ErrorField) (events.APIGatewayProxyResponse, error) {
	errorResponse := &ErrorResponse{
		Fields: fields,
	}
	return NewErrorResponse(statusCode, errorResponse)
}

func NewErrorResponse(statusCode int, errorResponse *ErrorResponse) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(errorResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
	}, nil
}
