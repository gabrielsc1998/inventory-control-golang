package id

import "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"

type ID struct {
	value string
}

func NewId(id string) (*ID, error) {
	if len(id) == 0 {
		return nil, custom_errors.InvalidParamError("id")
	}
	return &ID{value: id}, nil
}

func (id *ID) Value() string {
	return id.value
}
