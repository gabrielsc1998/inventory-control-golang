package custom_errors_test

import (
	"errors"

	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"
)

func TestShouldReceiveAExpectedError_InvalidParam(t *testing.T) {
	error := custom_errors.InvalidParamError("param")

	const expectedText = custom_errors.InvalidParam + " param"

	assert.Equal(t, error, errors.New(expectedText))
	assert.Equal(t, error.Error(), expectedText)
}
