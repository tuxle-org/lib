package entities_test

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/tuxle-org/lib/tuxle/entities"
	"gotest.tools/assert"
)

const SERVER_SERIALIZED = "server:v1\x00\x00\x00\x00\x00\x00\x00\x00EReal Name\x00Actual Description.\x00/tmp/nothing.jpg\x00"

func server() entities.Entity {
	return &entities.Server{
		ServerId:    69,
		Name:        "Real Name",
		Description: "Actual Description.",
		IconURI:     "/tmp/nothing.jpg",
	}
}

func TestSerialize(test *testing.T) {
	server := server()

	var buffer bytes.Buffer
	err := server.Serialize(&buffer)
	assert.NilError(test, err)

	assert.DeepEqual(
		test,
		buffer.String(),
		SERVER_SERIALIZED,
	)
}

func TestDeserialize(test *testing.T) {
	var out = new(entities.Server)
	reader := bufio.NewReader(strings.NewReader(SERVER_SERIALIZED))

	err := out.Deserialize(reader)
	assert.NilError(test, err)
	assert.DeepEqual(test, out, server())

	_, err = reader.ReadByte()
	assert.ErrorType(test, err, io.EOF)
}
