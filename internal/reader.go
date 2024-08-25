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

func (reader *Reader) ReadByte() (byte, error) {
	var data = make([]byte, 1)
	_, err := reader.IO.Read(data)
	if err != nil {
		return 0, fmt.Errorf("Error reading byte: %w", err)
	}

	return data[0], nil
}

func (reader *Reader) ReadRunes(field string, data []rune) error {
	var char rune
	for i := range data {
		if err := binary.Read(reader.IO, binary.LittleEndian, &char); err != nil {
			return fmt.Errorf("Error reading: %q: %w", field, err)
		}
		data[i] = char
	}

	return nil
}

func (reader *Reader) ReadString(field string, delimeter byte) (string, error) {
	var buffer strings.Builder

	for {
		char, err := reader.ReadByte()
		if err != nil {
			return "", err
		}
		if char == delimeter {
			break
		}
		buffer.WriteByte(char)
	}

	return buffer.String(), nil
}

func (reader *Reader) ReadUint32(field string) (uint32, error) {
	var data = make([]byte, 4)
	err := reader.Read(field, data)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint32(data), nil
}
