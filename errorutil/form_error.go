package errorutil

import (
	"encoding/json"
)

// FormError defines form error
type FormError struct {
	Message string        `json:"message,omitempty"`
	Fields  []*FieldError `json:"fields,omitempty"`
}

// NewFormError creates form error
func NewFormError(message string, fields []*FieldError) *FormError {
	if message == "" || fields == nil || len(fields) == 0 {
		return nil
	}

	return &FormError{
		Message: message,
		Fields:  fields,
	}
}

func (e *FormError) Error() string {
	field, err := json.Marshal(&FormError{
		Message: e.Message,
		Fields:  e.Fields,
	})
	if err != nil {
		return err.Error()
	}

	return string(field)
}

// FieldError defines field error
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
