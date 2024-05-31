package model

import (
	"time"
)

type ApiError struct {
	Code    int          `json:"code"`
	Type    string       `json:"type"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(code int, errortype, message string, data interface{}) *ApiError {
	err := &ApiError{
		Code:    code,
		Type:    errortype,
		Message: message,
		Data:    &data,
	}
	return err
}

type ApiMeta struct {
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *ApiError   `json:"error"`
	Meta    ApiMeta     `json:"meta"`
}
