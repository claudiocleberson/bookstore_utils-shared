package rest_err

import (
	"errors"
	"fmt"
	"net/http"
)

//Interface definition
type RestErr interface {
	Message() string
	Code() int
	Error() string
	Causes() []interface{}
}

//Inteface implementation
type restErr struct {
	message string        `json:"message"`
	code    int           `json:"code"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Code() int {
	return e.code
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

func (r restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]",
		r.message, r.code, r.error, r.causes)
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusBadRequest,
		error:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusNotFound,
		error:   "Not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	returnErr := restErr{
		message: message,
		code:    http.StatusInternalServerError,
		error:   "internal_server_error",
	}

	if err != nil {
		returnErr.causes = append(returnErr.causes, err.Error())
	}
	return returnErr
}

func NewRestError(message string, code int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		code:    code,
		error:   err,
		causes:  causes,
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		code:    http.StatusUnauthorized,
		error:   "not_authorized",
	}
}
func NewError(msg string) error {
	return errors.New(msg)
}
