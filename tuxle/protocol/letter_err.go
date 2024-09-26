package protocol

import (
	"io"

	"github.com/bbfh-dev/alt/alt"
)

const IdErrLetter = 1

type ErrLetter struct {
	Body string
}

func (letter ErrLetter) Read(reader io.Reader) (Letter, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, alt.PrefixErr("Reading Letter.Body", err)
	}

	letter.Body = string(data)

	return letter, nil
}

func (letter ErrLetter) Write(writer io.Writer) error {
	_, err := writer.Write([]byte(letter.Body))
	return alt.PrefixErr("Writing Letter.Body", err)
}
