package errors

import "fmt"

type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Cause      error
}

func (e *AppError) Error() string {
	if e.Cause == nil {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
}

func (e *AppError) Unwrap() error { return e.Cause }

func ValidationError(msg string) *AppError { return &AppError{Code: "validation_error", Message: msg, StatusCode: 400} }
func NotFoundError(msg string) *AppError   { return &AppError{Code: "not_found", Message: msg, StatusCode: 404} }
func UnauthorizedError(msg string) *AppError {
	return &AppError{Code: "unauthorized", Message: msg, StatusCode: 401}
}
func InternalError(msg string, cause error) *AppError {
	return &AppError{Code: "internal_error", Message: msg, StatusCode: 500, Cause: cause}
}
