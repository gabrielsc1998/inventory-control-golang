package customerrors_test

import (
	"errors"

	"github.com/stretchr/testify/assert"

	"testing"

	customerrors "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/errors/invalid-param"
)

func TestShouldReceiveAExpectedError(t *testing.T) {
	error := customerrors.InvalidParamError("param")
	assert.Equal(t, error, errors.New("Invalid Param - param"))
	assert.Equal(t, error.Error(), "Invalid Param - param")
}
