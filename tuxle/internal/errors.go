package internal

import "fmt"

func AnyErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func PrefixErr(prefix string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", prefix, err)
}
