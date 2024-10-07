package protocol

import (
	"fmt"
	"io"

	"github.com/bbfh-dev/go-tools/tools/terr"
)

type Letter interface {
	Read(io.Reader) (Letter, error)
	Write(writer io.Writer) error
}

func LetterType(reader io.Reader) (Letter, error) {
	var binaryType = make([]byte, 1)
	_, err := reader.Read(binaryType)
	if err != nil {
		return nil, err
	}

	var out Letter = nil
	switch binaryType[0] {
	case IdOkayLetter:
		out = OkayLetter{}
	case IdErrLetter:
		out = ErrLetter{}
	case IdEntityLetter:
		out = EntityLetter{}
	}

	if out == nil {
		return nil, fmt.Errorf("Unrecognized letter type of index: %d", binaryType[0])
	}

	return out, nil
}

func ReadLetter(reader io.Reader) (Letter, error) {
	out, err := LetterType(reader)
	if err != nil {
		return nil, err
	}
	return out.Read(reader)
}

func WriteLetter(letter Letter, writer io.Writer) error {
	var id uint8
	switch letter.(type) {
	case OkayLetter:
		id = IdOkayLetter
	case ErrLetter:
		id = IdErrLetter
	case EntityLetter:
		id = IdEntityLetter
	}

	_, err := writer.Write([]byte{id})
	if err != nil {
		return terr.Prefix("Writing Letter type identifier", err)
	}

	return letter.Write(writer)
}
