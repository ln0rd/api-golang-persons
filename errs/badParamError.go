package errs

import (
	"fmt"
	"net/http"
)

type BadParamError struct {
	ErrorName string `json:"error_name"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func NewBadParamError(message string) *BadParamError {
	return &BadParamError{ErrorName: "BadParamError", Code: http.StatusBadRequest, Message: message}
}

// Implementation of the error interface
func (e *BadParamError) Error() string {
	return fmt.Sprintf("%s: %s (Status: %d)", e.ErrorName, e.Message, e.Code)
}
