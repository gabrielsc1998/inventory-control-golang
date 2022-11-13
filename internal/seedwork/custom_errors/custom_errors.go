package custom_errors

import "errors"

func InvalidParamError(param string) error {
	return errors.New(InvalidParam + param)
}
