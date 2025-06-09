package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ln0rd/api-golang-persons/errs"
)

type ErrorResponse struct {
	ErrorName string `json:"error_name"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func HandleControllerError(w http.ResponseWriter, err error) {
	if customErr, ok := err.(*errs.CustomError); ok {
		w.WriteHeader(customErr.Code)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorName: customErr.ErrorName,
			Code:      customErr.Code,
			Message:   customErr.Message,
		})
		return
	}

	if notFoundErr, ok := err.(*errs.NotFoundError); ok {
		w.WriteHeader(notFoundErr.Code)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorName: notFoundErr.ErrorName,
			Code:      notFoundErr.Code,
			Message:   notFoundErr.Message,
		})
		return
	}

	if validateRequestErr, ok := err.(*errs.ValidateRequestError); ok {
		w.WriteHeader(validateRequestErr.Code)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorName: validateRequestErr.ErrorName,
			Code:      validateRequestErr.Code,
			Message:   validateRequestErr.Message,
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{
		ErrorName: "Internal Server Error",
		Code:      http.StatusInternalServerError,
		Message:   "Internal Server Error",
	})
}
