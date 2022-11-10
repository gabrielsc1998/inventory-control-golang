package customerror_test

import (
	"errors"

	"github.com/stretchr/testify/assert"

	"testing"

	customerror "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/errors"
)

func TestShouldReceiveAExpectedError_InvalidParam(t *testing.T) {
	error := customerror.InvalidParamError("param")
	assert.Equal(t, error, errors.New("Invalid Param - param"))
	assert.Equal(t, error.Error(), "Invalid Param - param")
}
