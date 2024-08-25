package field_test

import (
	"bytes"
	"testing"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/tuxle/field"
	"gotest.tools/assert"
)

func mockParameters() field.Parameters {
	return field.Parameters{
		"Hello":           "World",
		"contains spaces": "12345",
		"Привет Мир!":     "hi=world=123",
		"body!":           "header\nbody\nfooter",
	}
}

func mockParametersBuffer(test *testing.T) *bytes.Buffer {
	var buffer bytes.Buffer

	writer := internal.NewWriter(&buffer)
	err := internal.AnyErr(
		writer.Write(
			"Parameters",
			[]byte(
				"Hello=World\x00body!=header\nbody\nfooter\x00contains spaces=12345\x00Привет Мир!=hi=world=123\x00",
			),
		),
	)
	if err != nil {
		test.Fatal(err)
	}

	return &buffer
}

func TestParametersWrite(test *testing.T) {
	var buffer bytes.Buffer

	err := mockParameters().Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, buffer.Bytes(), mockParametersBuffer(test).Bytes())
}

func TestParametersRead(test *testing.T) {
	params := field.Parameters{}
	err := params.ReadUntilEOF(internal.NewReader(mockParametersBuffer(test)))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, params, mockParameters())
}

func TestParametersValidator(test *testing.T) {
	params := field.Parameters{"B": "hello", "C": ""}
	err := params.Validate(map[string]field.Validator{"A": field.Exists})
	if err == nil {
		test.Fatalf("Params must be invalid")
	}

	err = params.Validate(map[string]field.Validator{"B": field.Exists})
	if err != nil {
		test.Fatal(err)
	}

	err = params.Validate(map[string]field.Validator{"C": field.NotEmpty})
	if err == nil {
		test.Fatalf("Params must be invalid")
	}

	err = params.Validate(map[string]field.Validator{"B": field.HasLength(5)})
	if err != nil {
		test.Fatal(err)
	}

	err = params.Validate(map[string]field.Validator{"B": field.HasLength(6)})
	if err == nil {
		test.Fatalf("Params must be invalid")
	}
}
