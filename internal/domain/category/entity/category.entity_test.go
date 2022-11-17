package entity_test

// import (
// 	"github.com/stretchr/testify/assert"

// 	"testing"

// 	entity_category "github.com/gabrielsc1998/inventory-control-golang/internal/domain/category/entity"
// 	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"
// )

// func TestShouldCreateACategory(t *testing.T) {
// 	const inputId = "valid-id"
// 	const inputName = "valid-name"

// 	category, error := entity_category.NewCategory(inputId, inputName)

// 	assert.Equal(t, error, nil)
// 	assert.Equal(t, category.Id(), inputId)
// 	assert.Equal(t, category.Name(), inputName)
// 	assert.True(t, category.IsActive())
// }

// func TestShouldChangeCategoryName(t *testing.T) {
// 	const inputId = "valid-id"
// 	const inputName = "valid-name"

// 	category, error := entity_category.NewCategory(inputId, inputName)

// 	assert.Equal(t, error, nil)
// 	assert.Equal(t, category.Id(), inputId)
// 	assert.Equal(t, category.Name(), inputName)
// 	assert.True(t, category.IsActive())

// 	const inputNewName = "new-name"
// 	category.ChangeName(inputNewName)
// 	assert.Equal(t, category.Name(), inputNewName)
// }

// func TestShouldReturnErrorWhenTheIdIsInvalid(t *testing.T) {
// 	category, error := entity_category.NewCategory("", "valid-name")
// 	assert.Nil(t, category)
// 	assert.Equal(t, error, custom_errors.InvalidParamError("id"))
// }

// func TestShouldReturnErrorWhenTheNameIsInvalid(t *testing.T) {
// 	category, error := entity_category.NewCategory("valid-id", "")
// 	assert.Nil(t, category)
// 	assert.Equal(t, error, custom_errors.InvalidParamError("name"))
// }
