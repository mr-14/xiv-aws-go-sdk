package httputil

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/mr-14/xiv-aws-go-sdk/errorutil"
)

// NewMessageError creates message error
func NewMessageError(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       message,
	}, nil
}

// NewFieldError creates field error
func NewFieldError(statusCode int, fields []*errorutil.FieldError) (events.APIGatewayProxyResponse, error) {
	formError := &errorutil.FormError{
		Fields: fields,
	}
	return NewFormError(statusCode, formError)
}

// NewFormError creates error response
func NewFormError(statusCode int, formError *errorutil.FormError) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       formError.Error(),
	}, nil
}

// NewErrorResponse creates error response
func NewErrorResponse(err error) (events.APIGatewayProxyResponse, error) {
	switch e := err.(type) {
	case *errorutil.HTTPError:
		return events.APIGatewayProxyResponse{
			StatusCode: e.Status,
			Body:       e.Form.Error(),
		}, nil
	default:
		return events.APIGatewayProxyResponse{}, err
	}
}
