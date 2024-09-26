package protocol

import (
	"io"

	"github.com/tuxle-org/lib/tuxle/internal"
)

const IdErrLetter = 1

type ErrLetter struct {
	Body string
}

func (letter ErrLetter) Read(reader io.Reader) (Letter, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, internal.PrefixErr("Reading Letter.Body", err)
	}

	letter.Body = string(data)

	return letter, nil
}

func (letter ErrLetter) Write(writer io.Writer) error {
	_, err := writer.Write([]byte(letter.Body))
	return internal.PrefixErr("Writing Letter.Body", err)
}
