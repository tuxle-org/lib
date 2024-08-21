package tuxle

import (
	"fmt"
	"strings"

	"github.com/tuxle-org/lib/tuxle/security"
)

const USER_BYTE_SIZE = 16 + (32 * 4) + 64 + 256 + (security.SALT_LENGTH + security.HASH_LENGTH)

type User struct {
	Id          [16]byte
	Name        [32]rune
	PictureId   [64]byte
	Description *[256]byte
	Password    security.Password
}

func NewUser() *User {
	return &User{
		Id:          [16]byte{},
		Name:        [32]rune{},
		PictureId:   [64]byte{},
		Description: nil,
		Password:    security.Password{},
	}
}

// Ensures no pointers are equal to nil by setting them to empty arrays.
func (user *User) FillNilValues() *User {
	if user.Description == nil {
		user.Description = &[256]byte{}
	}

	return user
}

// WARN: NOT TESTED, since its contents shouldn't be relied on by any program.
// Use it only to display data.
func (user *User) String() string {
	var description [256]byte
	if user.Description != nil {
		description = *user.Description
	}
	return strings.ReplaceAll(fmt.Sprintf(
		"Id=%s\nName=%s\nPictureName=%s\nDescription=%s\nPassword=%s",
		user.Id,
		string(user.Name[:]),
		string(user.PictureId[:]),
		description,
		user.Password.String(),
	), "\x00", ".")
}
