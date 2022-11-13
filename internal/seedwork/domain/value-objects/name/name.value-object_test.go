package name_test

import (
	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"
	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/name"
)

func TestShouldCreateAName(t *testing.T) {
	const inputName = "valid-name"
	name, error := name.NewName(inputName)
	assert.Equal(t, name.Value(), inputName)
	assert.Equal(t, error, nil)
}

func TestShouldReturnErrorWhenTheIDIsInvalid(t *testing.T) {
	name, error := name.NewName("")
	assert.Nil(t, name)
	assert.Equal(t, error, custom_errors.InvalidParamError("name"))
}
