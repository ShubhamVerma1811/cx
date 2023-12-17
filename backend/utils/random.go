package utils

import (
	"math/rand"
)

func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	str := make([]byte, n)

	for i := range str {
		str[i] = byte(letters[rand.Intn(len(letters))])
	}
	return string(str)
}
