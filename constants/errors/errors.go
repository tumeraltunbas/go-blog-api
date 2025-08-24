package errors

import "net/http"

// Error Codes
const ProcessFailureError string = "PROCESS_FAILURE_ERROR"
const AuthorizationError string = "AUTHORIZATION_ERROR"

// Error Messages
var ErrorMessages = map[string]string{
	ProcessFailureError: "Proceess failure error.",
	AuthorizationError:  "Authorization error.",
}

type BaseError struct {
	Code    string
	Message string
	Status  int
}

func NewInternalServerError() BaseError {
	err := BaseError{Code: ProcessFailureError, Message: ErrorMessages[ProcessFailureError], Status: http.StatusInternalServerError}
	return err
}

func NewBusinessRuleError(code string) BaseError {
	err := BaseError{Code: code, Message: ErrorMessages[code], Status: http.StatusUnprocessableEntity}
	return err
}

func NewAuthorizationError() BaseError {
	err := BaseError{Code: AuthorizationError, Message: ErrorMessages[AuthorizationError], Status: http.StatusUnauthorized}
	return err
}
