package rest_response

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"errors"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	//TODO: Test!
}

func TestNewNotFoundError(t *testing.T) {
	//TODO: Test!
}

func TestNewError(t *testing.T) {
	//TODO: Test!
}
