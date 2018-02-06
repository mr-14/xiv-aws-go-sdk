package exception

import "net/http"

// HTTPError defines HTTP error
type HTTPError struct {
	Status int
	Code   string
	Fields []*FieldError
}

func (e *HTTPError) Error() string {
	return e.Code
}

// NewBadRequest creates a 400 error
func NewBadRequest(code string) error {
	return &HTTPError{http.StatusBadRequest, code, nil}
}

// NewBadFormRequest creates a 400 error with list of form errors
func NewBadFormRequest(code string, fields []*FieldError) error {
	return &HTTPError{http.StatusBadRequest, code, fields}
}

// NewUnauthorized creates a 401 error
func NewUnauthorized(code string) error {
	return &HTTPError{http.StatusUnauthorized, code, nil}
}

// NewNotFound creates a 404 error
func NewNotFound(code string) error {
	return &HTTPError{http.StatusNotFound, code, nil}
}

// NewServiceUnavailable creates a 503 error
func NewServiceUnavailable(code string) error {
	return &HTTPError{http.StatusServiceUnavailable, code, nil}
}
