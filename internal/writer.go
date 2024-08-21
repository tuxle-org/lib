package internal

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: writer,
	}
}

func (writer *Writer) Write(field string, data []byte) error {
	_, err := writer.writer.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing: %q: %w", field, err)
	}

	return nil
}

func (writer *Writer) WriteRunes(field string, data []rune) error {
	for _, char := range data {
		if err := binary.Write(writer.writer, binary.LittleEndian, char); err != nil {
			return fmt.Errorf("Error writing: %q: %w", field, err)
		}
	}

	return nil
}

func (writer *Writer) WriteString(field string, str string, length int) error {
	data := make([]byte, length)
	copy(data[:], []byte(str))

	return writer.Write(field, data)
}
