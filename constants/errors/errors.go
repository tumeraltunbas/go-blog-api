package errors

import (
	"encoding/json"
	"net/http"
)

// Error Codes
const ProcessFailureError string = "PROCESS_FAILURE_ERROR"
const AuthorizationError string = "AUTHORIZATION_ERROR"
const UserAlreadyExist string = "USER_ALREADY_EXIST"
const ValidationError string = "VALIDATION_ERROR"

// Error Messages
var ErrorMessages = map[string]string{
	ProcessFailureError: "Proceess failure error.",
	AuthorizationError:  "Authorization error.",
	UserAlreadyExist:    "User already exist.",
}

type BaseError struct {
	Code    string
	Message string
	Status  int
}

func handleError(baseError BaseError, w http.ResponseWriter) {
	errorObject := map[string]string{
		"Code":    baseError.Code,
		"Message": baseError.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(baseError.Status)
	json.NewEncoder(w).Encode(errorObject)
}

func NewInternalServerError(w http.ResponseWriter) {
	err := BaseError{Code: ProcessFailureError, Message: ErrorMessages[ProcessFailureError], Status: http.StatusInternalServerError}
	handleError(err, w)
}

func NewBusinessRuleError(code string, w http.ResponseWriter) {
	err := BaseError{Code: code, Message: ErrorMessages[code], Status: http.StatusUnprocessableEntity}
	handleError(err, w)
}

func NewAuthorizationError(w http.ResponseWriter) {
	err := BaseError{Code: AuthorizationError, Message: ErrorMessages[AuthorizationError], Status: http.StatusUnauthorized}
	handleError(err, w)
}

func NewValidationError(w http.ResponseWriter, data error) {
	err := BaseError{Code: ValidationError, Message: data.Error(), Status: http.StatusBadRequest}
	handleError(err, w)
}
