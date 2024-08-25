package stream

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Writer struct {
	IO io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		IO: writer,
	}
}

func (writer *Writer) Write(field string, data []byte) error {
	_, err := writer.IO.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing: %q: %w", field, err)
	}

	return nil
}

func (writer *Writer) WriteRunes(field string, data []rune) error {
	for _, char := range data {
		if err := binary.Write(writer.IO, binary.LittleEndian, char); err != nil {
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

func (writer *Writer) WriteUint32(field string, number uint32) error {
	var buffer = make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, number)
	return writer.Write(field, buffer)
}
