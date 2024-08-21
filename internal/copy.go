package internal

func CopyBytes(str string, length int) []byte {
	data := make([]byte, length)
	copy(data[:], []byte(str))

	return data
}

func CopyRunes(str string, length int) []rune {
	data := make([]rune, length)
	copy(data[:], []rune(str))

	return data
}
