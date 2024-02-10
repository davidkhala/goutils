package http

import (
	"net/http"
)

type Error struct {
	Code    StatusCode
	Message string
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return "[" + e.Code.String() + "]" + e.Message
}

func BadRequest(msg string) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func ServiceUnavailable(msg string) Error {
	return Error{
		Code:    http.StatusServiceUnavailable,
		Message: msg,
	}
}
func InternalServerError(msg string) Error {
	return Error{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
