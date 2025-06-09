package errs

import "fmt"

type CustomError struct {
	ErrorName string `json:"error_name"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func NewCustomError(name string, statusCode int, message string) *CustomError {
	return &CustomError{
		ErrorName: name,
		Code:      statusCode,
		Message:   message,
	}
}

// Implementation of the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s (Status: %d)", e.ErrorName, e.Message, e.Code)
}
