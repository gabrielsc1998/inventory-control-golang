package id_test

import (
	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"

	valueobject_id "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/id"
)

func TestShouldCreateAnID(t *testing.T) {
	const inputId = "valid-id"
	id, error := valueobject_id.NewId(inputId)
	assert.Equal(t, id.Value(), inputId)
	assert.Equal(t, error, nil)
}

func TestShouldReturnErrorWhenTheIDIsInvalid(t *testing.T) {
	id, error := valueobject_id.NewId("")
	assert.Nil(t, id)
	assert.Equal(t, error, custom_errors.InvalidParamError("- id"))
}
