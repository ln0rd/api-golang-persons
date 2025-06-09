package errs

import (
	"fmt"
	"net/http"
)

type ValidateRequestError struct {
	ErrorName string `json:"error_name"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func NewValidateRequestError(message string) *ValidateRequestError {
	return &ValidateRequestError{ErrorName: "ValidateRequestError", Code: http.StatusBadRequest, Message: message}
}

// Implementation of the error interface
func (e *ValidateRequestError) Error() string {
	return fmt.Sprintf("%s: %s (Status: %d)", e.ErrorName, e.Message, e.Code)
}
