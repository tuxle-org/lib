package protocol

import (
	"encoding/gob"
	"io"
)

const IdEntityLetter = 2

type EntityLetter struct {
	Entity interface{}
}

func (letter EntityLetter) Read(reader io.Reader) (Letter, error) {
	err := gob.NewDecoder(reader).Decode(&letter.Entity)
	return letter, err
}

func (letter EntityLetter) Write(writer io.Writer) error {
	return gob.NewEncoder(writer).Encode(letter.Entity)
}
