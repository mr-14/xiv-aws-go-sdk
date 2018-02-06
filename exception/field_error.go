package exception

// FieldError defines field error
type FieldError struct {
	Field string `json:"field"`
	Code  string `json:"code"`
}

// NewFieldError creates a field error
func NewFieldError(name string, code string) *FieldError {
	return &FieldError{name, code}
}
