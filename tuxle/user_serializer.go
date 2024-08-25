package tuxle

import (
	"github.com/tuxle-org/lib/internal"
	"github.com/tuxle-org/lib/stream"
)

func (user *User) Size() int64 {
	return USER_BYTE_SIZE
}

func (user *User) Serialize(writer *stream.Writer) error {
	user.FillNilValues()

	return internal.AnyErr(
		writer.Write("user.Id", user.Id[:]),
		writer.WriteRunes("user.Name", user.Name[:]),
		writer.Write("user.PictureId", user.PictureId[:]),
		writer.Write("user.Description", user.Description[:]),
		user.Password.Write(writer),
	)
}

func (user *User) Deserialize(reader *stream.Reader) error {
	user.FillNilValues()

	return internal.AnyErr(
		reader.Read("user.Id", user.Id[:]),
		reader.ReadRunes("user.Name", user.Name[:]),
		reader.Read("user.PictureId", user.PictureId[:]),
		reader.Read("user.Description", user.Description[:]),
		user.Password.Read(reader),
	)
}
