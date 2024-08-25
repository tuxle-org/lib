package field

import (
	"errors"
	"fmt"
)

type Validator func(string) error

func Exists(_ string) error {
	return nil
}

func HasLength(length int) Validator {
	return func(value string) error {
		if len(value) != length {
			return fmt.Errorf("Field must be `%d` characters long, not `%d`", length, len(value))
		}

		return nil
	}
}

func NotEmpty(value string) error {
	if len(value) == 0 {
		return errors.New("Field cannot be empty")
	}

	return nil
}
