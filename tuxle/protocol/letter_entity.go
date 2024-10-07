package protocol

import (
	"bytes"
	"encoding/gob"
	"io"

	"github.com/bbfh-dev/go-tools/tools/terr"
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

func EncodeEntity(entity interface{}) (EntityLetter, error) {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(entity)
	if err != nil {
		return EntityLetter{}, terr.Prefix("Encoding entity", err)
	}

	return EntityLetter{Entity: buffer.Bytes()}, nil
}

func (letter EntityLetter) DecodeEntity(target interface{}) error {
	return gob.NewDecoder(bytes.NewReader(letter.Entity)).Decode(target)
}
