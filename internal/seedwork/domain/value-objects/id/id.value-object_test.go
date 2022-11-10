package valueobject_id_test

import (
	"github.com/stretchr/testify/assert"

	"testing"

	valueobject_id "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/value-objects/id"
	customerror "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/errors"
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
	assert.Equal(t, error, customerror.InvalidParamError("id"))
}
