package protocol_test

import (
	"bytes"
	"testing"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/stream"
	"github.com/tuxle-org/lib/tuxle/field"
	"github.com/tuxle-org/lib/tuxle/protocol"
	"gotest.tools/assert"
)

func mockLetter() *protocol.Letter {
	return &protocol.Letter{
		Type:     "TEST",
		Endpoint: "example",
		Parameters: field.Parameters{
			"Key": "Value",
		},
		Body: "Hello World!\nПривет Мир!",
	}
}

func mockLetterBuffer(test *testing.T) *bytes.Buffer {
	var buffer bytes.Buffer

	writer := stream.NewWriter(&buffer)
	err := internal.AnyErr(
		writer.Write("Letter", []byte("TEST example\n")),
		writer.WriteUint32("Parameters.@len", 1),
		writer.Write("Letter", []byte("Key=Value\x00Hello World!\nПривет Мир!\r")),
	)
	if err != nil {
		test.Fatal(err)
	}

	return &buffer
}

func TestLetterWrite(test *testing.T) {
	var buffer bytes.Buffer

	err := mockLetter().Write(stream.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, buffer.Bytes(), mockLetterBuffer(test).Bytes())
}

func TestLetterRead(test *testing.T) {
	letter := protocol.NewLetter()
	err := letter.Read(stream.NewReader(mockLetterBuffer(test)))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, letter, mockLetter())
}
