package protocol

import (
	"fmt"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/stream"
	"github.com/tuxle-org/lib/tuxle/field"
)

// A letter is used to communicate between programs.
type Letter struct {
	Type       string
	Endpoint   string
	Parameters field.Parameters
	Body       string
}

// Get an empty letter
func NewLetter() *Letter {
	return &Letter{
		Type:       "",
		Endpoint:   "",
		Parameters: field.Parameters{},
		Body:       "",
	}
}

func (letter *Letter) Write(writer *stream.Writer) error {
	_, err := fmt.Fprintf(writer.IO, "%s %s\n", letter.Type, letter.Endpoint)
	if err != nil {
		return err
	}

	err = internal.AnyErr(
		writer.WriteUint32("Letter.Parameters.@len", uint32(len(letter.Parameters))),
		letter.Parameters.Write(writer),
	)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer.IO, "%s\r", letter.Body)
	return err
}

func (letter *Letter) Read(reader *stream.Reader) (err error) {
	letter.Type, err = reader.ReadString("Letter.Type", ' ')
	if err != nil {
		return err
	}

	letter.Endpoint, err = reader.ReadString("Letter.Endpoint", '\n')
	if err != nil {
		return err
	}

	count, err := reader.ReadUint32("Letter.Parameters.@len")
	if err != nil {
		return err
	}

	err = letter.Parameters.Read(reader, count)
	if err != nil {
		return err
	}

	letter.Body, err = reader.ReadString("Letter.Body", '\r')

	return err
}
