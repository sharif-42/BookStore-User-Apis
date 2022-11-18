package errors

import (
	"net/http"
)

type RestError struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	ErrorCode string `json:"error_code"`
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Message:   message,
		Status:    http.StatusBadRequest,
		ErrorCode: "Bad Request",
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message:   message,
		Status:    http.StatusNotFound,
		ErrorCode: "Not Found!",
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Message:   message,
		Status:    http.StatusInternalServerError,
		ErrorCode: "Internal Server Error!",
	}
}
