package protocol_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/tuxle-org/lib/tuxle/protocol"
	"gotest.tools/assert"
)

func isType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func TestTypes(test *testing.T) {
	var letters = map[uint8]protocol.Letter{
		protocol.IdOkayLetter: protocol.OkayLetter{},
		protocol.IdErrLetter:  protocol.ErrLetter{},
	}

	for letterIndex, letter := range letters {
		var buffer bytes.Buffer
		buffer.Write([]byte{letterIndex})

		recognizedLetter, err := protocol.ReadLetter(&buffer)
		assert.NilError(test, err)

		if !isType(recognizedLetter, letter) {
			test.Fatalf("Invalid letter type, must be %q", reflect.TypeOf(letter))
		}
		test.Log("Recognized", letterIndex, "->", reflect.TypeOf(letter))
	}
}

func writeLetter(test *testing.T, letter protocol.Letter) bytes.Buffer {
	var buffer bytes.Buffer
	err := protocol.WriteLetter(letter, &buffer)
	if err != nil {
		test.Fatalf("%s: %s", reflect.TypeOf(letter), err.Error())
	}
	return buffer
}

func TestWrite(test *testing.T) {
	var letters = map[protocol.Letter]string{
		protocol.OkayLetter{}:                    "",
		protocol.ErrLetter{Body: "Hello World!"}: "Hello World!",
	}

	for letter, body := range letters {
		buffer := writeLetter(test, letter)

		assert.Equal(test, string(buffer.Bytes()[1:]), body)
	}
}

func TestRead(test *testing.T) {
	var letters = []protocol.Letter{
		protocol.OkayLetter{},
		protocol.ErrLetter{Body: "Hello World!"},
	}

	for _, letter := range letters {
		buffer := writeLetter(test, letter)

		out, err := protocol.ReadLetter(&buffer)
		if err != nil {
			test.Fatalf(
				"%s: %s (Recognized as %s)",
				reflect.TypeOf(letter),
				err.Error(),
				reflect.TypeOf(out),
			)
		}

		assert.Equal(test, out, letter)
	}
}
