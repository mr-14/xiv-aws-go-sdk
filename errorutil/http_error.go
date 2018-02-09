package errorutil

import (
	"encoding/json"
)

// HTTPError defines form error
type HTTPError struct {
	Status int        `json:"status"`
	Form   *FormError `json:"form,omitempty"`
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
