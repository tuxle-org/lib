package tuxle_test

import (
	"bytes"
	"testing"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/tuxle"
	"github.com/tuxle-org/lib/tuxle/field"
	"gotest.tools/assert"
)

func mockServer() *tuxle.Server {
	server := tuxle.NewServer()
	server.Name = "Test"
	server.Description = "# Hello world!\nWelcome to my server."
	return server
}

func mockServerBuffer(test *testing.T) *bytes.Buffer {
	var buffer bytes.Buffer

	err := field.Parameters{
		"Name":        "Test",
		"Description": "# Hello world!\nWelcome to my server.",
		"PictureId":   "",
	}.Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	return &buffer
}

func mockFullServer() *tuxle.Server {
	server := mockServer()
	server.PictureId = "YTEyamszbmtzamJhbjAxMm5rbHNqbmFvcDFuMjMwbmFzCYTEyamszbmtzamJhbjA"
	return server
}

func mockFullServerBuffer(test *testing.T) *bytes.Buffer {
	var buffer bytes.Buffer

	err := field.Parameters{
		"Name":        "Test",
		"Description": "# Hello world!\nWelcome to my server.",
		"PictureId":   "YTEyamszbmtzamJhbjAxMm5rbHNqbmFvcDFuMjMwbmFzCYTEyamszbmtzamJhbjA",
	}.Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	return &buffer
}

func TestServerWrite(test *testing.T) {
	var buffer bytes.Buffer
	err := mockServer().Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, buffer.Bytes(), mockServerBuffer(test).Bytes())
}

func TestFullServerWrite(test *testing.T) {
	var buffer bytes.Buffer
	err := mockFullServer().Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, buffer.Bytes(), mockFullServerBuffer(test).Bytes())
}

func TestServerRead(test *testing.T) {
	server := tuxle.NewServer()
	err := server.Read(internal.NewReader(mockServerBuffer(test)))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, server, mockServer())
}

func TestFullServerRead(test *testing.T) {
	server := tuxle.NewServer()
	err := server.Read(internal.NewReader(mockFullServerBuffer(test)))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, server, mockFullServer())
}
