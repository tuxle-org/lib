package field

import (
	"errors"
	"fmt"
	"io"
	"math"
	"sort"

	"github.com/tuxle-org/lib/stream"
	"golang.org/x/exp/maps"
)

type Parameters map[string]string

func (params Parameters) Keys() (keys []string) {
	keys = maps.Keys(params)
	sort.Strings(keys)
	return
}

func (params Parameters) Write(writer *stream.Writer) error {
	var err error
	for _, key := range params.Keys() {
		_, err = fmt.Fprintf(writer.IO, "%s=%s\x00", key, params[key])
		if err != nil {
			return err
		}
	}

	return nil
}

func (params Parameters) Read(reader *stream.Reader, count uint32) error {
	for range count {
		key, err := reader.ReadString("@parameters:Key", '=')
		if err != nil {
			return err
		}

		value, err := reader.ReadString("@parameters:Value", '\x00')
		if err != nil {
			return err
		}

		params[key] = value
	}

	return nil
}

func (params Parameters) ReadUntilEOF(reader *stream.Reader) error {
	err := params.Read(reader, math.MaxInt32)
	if errors.Is(err, io.EOF) {
		return nil
	}
	return err
}

// Ensures all the fields are:
//
// 1. Defined.
//
// 2. Pass the Validator function without any errors.
//
// Returns the combined error of all fields that failed the check.
// Use `field.Exist()` Validator to only check that the field is present.
func (params Parameters) Validate(fields map[string]Validator) error {
	var errs []error

	for key, validator := range fields {
		value, ok := params[key]
		if !ok {
			errs = append(errs, fmt.Errorf("Field %q doesn't exist.", key))
			continue
		}
		if err := validator(value); err != nil {
			errs = append(errs, fmt.Errorf("Field %q validation failed: %w", key, err))
		}
	}

	return errors.Join(errs...)
}
