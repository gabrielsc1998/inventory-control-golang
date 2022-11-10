package valueobject_name_test

import (
	"github.com/stretchr/testify/assert"

	"testing"

	valueobject_name "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/name"
	customerror "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/errors"
)

func TestShouldCreateAName(t *testing.T) {
	const inputName = "valid-name"
	name, error := valueobject_name.NewName(inputName)
	assert.Equal(t, name.Value(), inputName)
	assert.Equal(t, error, nil)
}

func TestShouldReturnErrorWhenTheIDIsInvalid(t *testing.T) {
	name, error := valueobject_name.NewName("")
	assert.Nil(t, name)
	assert.Equal(t, error, customerror.InvalidParamError("name"))
}
