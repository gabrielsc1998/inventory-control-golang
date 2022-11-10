package customerrors

import "errors"

func InvalidParamError(param string) error {
	return errors.New("Invalid Param - " + param)
}
