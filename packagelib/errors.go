package packagelib

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func CekData(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "1" {
		return NotFoundError
	}

	return nil
}
