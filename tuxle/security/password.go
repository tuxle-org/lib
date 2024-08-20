package security

import (
	"github.com/tuxle-org/lib/internal"
)

const SALT_LENGTH = 8
const HASH_LENGTH = 64

type Password [SALT_LENGTH + HASH_LENGTH]byte

func (password Password) String() string {
	return string(password[:])
}

func (password *Password) Write(writer *internal.Writer) error {
	return internal.AnyErr(
		writer.Write("password:salt", password[:SALT_LENGTH]),
		writer.Write("password:hash", password[SALT_LENGTH:]),
	)
}

func (password *Password) Read(reader *internal.Reader) error {
	return internal.AnyErr(
		reader.Read("password:salt", password[:SALT_LENGTH]),
		reader.Read("password:hash", password[SALT_LENGTH:]),
	)
}

func (password Password) MatchesWith(str string) bool {
	var passwd Password
	copy(passwd[:SALT_LENGTH], password[:SALT_LENGTH])
	copy(passwd[SALT_LENGTH:], encodeString(string(password[:SALT_LENGTH])+str))

	return password == passwd
}
