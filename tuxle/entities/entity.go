package entities

import (
	"bufio"
	"io"
)

type Entity interface {
	Serialize(io.Writer) error
	Deserialize(*bufio.Reader) error
}
