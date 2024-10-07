package protocol

import (
	"bufio"
	"io"

	"github.com/bbfh-dev/go-tools/tools/tbin"
)

const IdGetLetter = 3

type GetLetter struct {
	Query query
}

func (letter GetLetter) Read(reader io.Reader) (Letter, error) {
	str, err := tbin.TuxleString(bufio.NewReader(reader))
	letter.Query = query(str)
	return letter, err
}

func (letter GetLetter) Write(writer io.Writer) error {
	return tbin.WriteTuxleString(writer, string(letter.Query))
}

type query string

const (
	GET_SERVER_INFO query = "server_info"
)
