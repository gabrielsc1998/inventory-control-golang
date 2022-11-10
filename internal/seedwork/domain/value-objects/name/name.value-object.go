package valueobject_name

import customerror "github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/errors"

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {
	if len(name) == 0 {
		return nil, customerror.InvalidParamError("name")
	}
	return &Name{value: name}, nil
}

func (name *Name) Value() string {
	return name.value
}
