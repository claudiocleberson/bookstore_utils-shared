package rest_err

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {

	err := NewInternalServerError("this is the message", errors.New("database_error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database_error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	//Todo: test!
}

func TestNewNotFoundError(t *testing.T) {
	//Todo: test
}

func TestNewError(t *testing.T) {
	//Todo - Test
}
