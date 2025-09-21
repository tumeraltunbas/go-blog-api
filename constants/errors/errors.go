package errors

import (
	"encoding/json"
	"net/http"
)

// Error Codes
const ProcessFailureError string = "PROCESS_FAILURE_ERROR"
const AuthorizationError string = "AUTHORIZATION_ERROR"
const UserNotFound string = "USER_NOT_FOUND"

// Error Messages
var ErrorMessages = map[string]string{
	ProcessFailureError: "Proceess failure error.",
	AuthorizationError:  "Authorization error.",
	UserNotFound:        "User not found.",
}

type BaseError struct {
	Code    string
	Message string
	Status  int
}

func handleError(error BaseError, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(error.Status)
	json.NewEncoder(w).Encode(error)
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
