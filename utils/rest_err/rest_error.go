package rest_err

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "Not_found",
	}
}

func NewInternalServerError(message string, err error) *RestErr {
	returnErr := &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	if err != nil {
		returnErr.Causes = append(returnErr.Causes, err.Error())
	}
	return returnErr
}

func NewRestError(message string, code int, err string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    code,
		Error:   err,
	}
}

func NewError(msg string) error {
	return errors.New(msg)
}
