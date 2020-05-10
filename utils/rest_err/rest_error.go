package rest_err

import (
	"encoding/json"
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
	ToJson() interface{}
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

func (r restErr) ToJson() interface{} {
	bytes, _ := json.Marshal(r)
	//res := restErr{}
	var dat map[string]interface{}
	err := json.Unmarshal(bytes, &dat)
	if err != nil {
		return nil
	}
	return dat
	//bytes, _ := json.Marshal(r)
	//return string(bytes)
	//return fmt.Sprintf(`"message": %s, "code": %d, "error": %s, "causes": [ %v ]`, r.message, r.code, r.error, r.causes)
	// return restErr{
	// 	message: r.message,
	// 	code:    r.code,
	// 	error:   r.error,
	// 	causes:  r.causes,
	// }
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, err
	}
	return apiErr, nil
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
