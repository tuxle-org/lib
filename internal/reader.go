package internal

import (
	"fmt"
	"io"
)

type Reader struct {
	reader io.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: reader,
	}
}

func (reader *Reader) Read(field string, data []byte) error {
	_, err := io.ReadFull(reader.reader, data)
	if err != nil {
		return fmt.Errorf("Error reading: %q: %w", field, err)
	}

	return nil
}
