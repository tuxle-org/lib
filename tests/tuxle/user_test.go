package tuxle_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/stream"
	"github.com/tuxle-org/lib/tuxle"
	"github.com/tuxle-org/lib/tuxle/security"
	"gotest.tools/assert"
)

func mockUser() *tuxle.User {
	user := tuxle.NewUser()
	user.Id = [16]byte(internal.CopyBytes("root", 16))
	user.Name = [32]rune(internal.CopyRunes("Root User", 32))
	user.PictureId = [64]byte(internal.CopyBytes("qwerty", 64))
	user.Description = nil
	user.Password = security.Password(
		internal.CopyBytes(
			"77e229855bbb7fa8d0837c224882756a55e41ba6670b6cf1a2fc696fa62acf347fc19c8c",
			security.HASH_LENGTH+security.SALT_LENGTH,
		),
	)
	return user
}

func mockUserBytes() string {
	return `root............R...o...o...t... ...U...s...e...r...............................................................................................qwerty..........................................................................................................................................................................................................................................................................................................................77e229855bbb7fa8d0837c224882756a55e41ba6670b6cf1a2fc696fa62acf347fc19c8c`
}

func TestUserSerialize(test *testing.T) {
	var buffer bytes.Buffer
	writer := stream.NewWriter(&buffer)

	user := mockUser()
	err := user.Serialize(writer)
	if err != nil {
		test.Fatal(err)
	}

	assert.Equal(
		test,
		strings.ReplaceAll(buffer.String(), "\x00", "."),
		mockUserBytes(),
	)
	assert.Equal(test, int64(buffer.Len()), user.Size())
}

func TestUserDeserialize(test *testing.T) {
	var buffer bytes.Buffer
	writer := stream.NewWriter(&buffer)

	user := mockUser()
	err := user.Serialize(writer)
	if err != nil {
		test.Fatal(err)
	}

	user = tuxle.NewUser()
	err = user.Deserialize(stream.NewReader(&buffer))
	if err != nil {
		test.Fatal(err)
	}

	assert.DeepEqual(test, user, mockUser().FillNilValues())
}

func TestUserString(test *testing.T) {
	if tuxle.NewUser().FillNilValues().Description == nil {
		test.Fatal("User.Description must not be nil!")
	}
	if len(tuxle.NewUser().String()) < 1 {
		test.Fatal("User.String() is empty!")
	}
}
