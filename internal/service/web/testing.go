package web

import "math/rand"

// fixedLengthString returns random string with fixed length.
func fixedLengthString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rs := make([]rune, length)
	for i := range rs {
		rs[i] = letters[rand.Intn(len(letters))]
	}

	return string(rs)
}
