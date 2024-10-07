package protocol

import (
	"io"
)

const IdEntityLetter = 2

type EntityLetter struct {
	Entity []byte
}

func (letter EntityLetter) Read(reader io.Reader) (Letter, error) {
	var err error
	letter.Entity, err = io.ReadAll(reader)
	return letter, err
}

func (letter EntityLetter) Write(writer io.Writer) error {
	_, err := writer.Write(letter.Entity)
	return err
}
