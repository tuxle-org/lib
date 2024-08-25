package internal

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type Reader struct {
	IO io.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		IO: reader,
	}
}

func (reader *Reader) Read(field string, data []byte) error {
	_, err := io.ReadFull(reader.IO, data)
	if err != nil {
		return fmt.Errorf("Error reading: %q: %w", field, err)
	}

	return nil
}

func (reader *Reader) ReadRunes(field string, data []rune) error {
	var char rune
	for i := range data {
		if err := binary.Read(reader.reader, binary.LittleEndian, &char); err != nil {
			return fmt.Errorf("Error reading: %q: %w", field, err)
		}
		data[i] = char
	}

	return nil
}
