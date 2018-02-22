package errorutil

import (
	"encoding/json"
)

// HTTPError defines form error
type HTTPError struct {
	Status int        `json:"status"`
	Form   *FormError `json:"form,omitempty"`
}

// NewHTTPMessageError creates HTTP error that consists of a single message
func NewHTTPMessageError(status int, message string) *HTTPError {
	return &HTTPError{
		Status: status,
		Form: &FormError{
			Message: message,
		},
	}
}

// NewHTTPFormError creates HTTP form error
func NewHTTPFormError(status int, form *FormError) *HTTPError {
	return &HTTPError{
		Status: status,
		Form:   form,
	}
}

func (e *HTTPError) Error() string {
	field, err := json.Marshal(HTTPError{
		Status: e.Status,
		Form:   e.Form,
	})
	if err != nil {
		return err.Error()
	}

	return string(field)
}
