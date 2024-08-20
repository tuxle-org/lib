package security_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/tuxle/security"
	"gotest.tools/assert"
)

func TestPassword(test *testing.T) {
	salt := strings.Repeat("0", security.SALT_LENGTH)
	hash := strings.Repeat("a", security.HASH_LENGTH)

	var passwd security.Password
	copy(passwd[:security.SALT_LENGTH], salt)
	copy(passwd[security.SALT_LENGTH:], hash)

	assert.DeepEqual(test, passwd, security.Password([]byte(salt+hash)))
}

func TestRandomHash(test *testing.T) {
	assert.Equal(test, len(security.RandomHash(128)), 128)
	assert.Equal(test, len(security.RandomHash(256)), 256)

	hash := security.RandomHash(256)
	if hash[:128] == hash[128:] {
		test.Fatalf("Hash should be random even past SHA512 limit!")
	}
}

func TestPasswordGen(test *testing.T) {
	str := "hello.world123"
	password := security.GenPassword(str)
	if !password.MatchesWith(str) {
		test.Fatalf("Password must match, since it was encoded from the same string.")
	}
}

func TestPasswordWrite(test *testing.T) {
	str := "hello.world123"
	password := security.GenPassword(str)

	var buffer bytes.Buffer
	err := password.Write(internal.NewWriter(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, password.String(), buffer.String())
}

func TestPasswordRead(test *testing.T) {
	str := "hello.world123"
	password := security.GenPassword(str)

	var buffer bytes.Buffer
	buffer.WriteString(password.String())

	passwd := security.Password{}
	err := passwd.Read(internal.NewReader(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, password, passwd)
}

func TestPasswordErrorHandling(test *testing.T) {
	var buffer [20]byte

	password := security.Password{}
	err := password.Read(internal.NewReader(bytes.NewReader(buffer[:])))
	if err == nil {
		test.Fatalf("Reading should fail, because there's not enough space!")
	}
}
