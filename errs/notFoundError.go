package errs

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	ErrorName string `json:"error_name"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{ErrorName: "NotFoundError", Code: http.StatusNotFound, Message: message}
}

// Implementation of the error interface
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s (Status: %d)", e.ErrorName, e.Message, e.Code)
}
