package protocol

import (
	"io"
)

const IdOkayLetter = 0

type OkayLetter struct {
	Body string
}

func (letter OkayLetter) Read(reader io.Reader) (Letter, error) {
	return letter, nil
}

func (letter OkayLetter) Write(writer io.Writer) error {
	return nil
}
