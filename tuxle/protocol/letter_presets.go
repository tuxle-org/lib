package protocol

import (
	"fmt"

	"github.com/tuxle-org/lib/tuxle/field"
)

func NewErrorLetter(name string, body string, args ...any) *Letter {
	return &Letter{
		Type:       "ERROR",
		Endpoint:   name,
		Parameters: field.Parameters{},
		Body:       fmt.Sprintf(body, args...),
	}
}

func NewOkayLetter(name string, body string, args ...any) *Letter {
	return &Letter{
		Type:       "OKAY",
		Endpoint:   name,
		Parameters: field.Parameters{},
		Body:       fmt.Sprintf(body, args...),
	}
}
