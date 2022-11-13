package name

import "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/custom_errors"

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	if len(name) == 0 {
		return nil, custom_errors.InvalidParamError("name")
	}
	return &Name{value: name}, nil
}

func (name *Name) Value() string {
	return name.value
}
