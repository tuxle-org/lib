package internal

import (
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
