package errs

import (
	"net/http"
	"yana-golang/logs"
)

type AppError struct {
	Code     int
	Message  string
	ErrorMsg interface{}
}

func (e AppError) Error() string {
	return e.Message
}

func ErrorNotFound(message string, any interface{}) error {
	logs.Error(any)
	return AppError{
		Code:     http.StatusNotFound,
		Message:  message,
		ErrorMsg: any,
	}
}

func ErrorBadRequest(message string, any interface{}) error {
	logs.Error(any)
	return AppError{
		Code:     http.StatusBadRequest,
		Message:  message,
		ErrorMsg: any,
	}
}

func ErrorInternalServer() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
}
