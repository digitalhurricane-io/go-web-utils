package utils

import (
	"math/rand"
)

var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var LowercaseLetters = []rune("abcdefghijklmnopqrstuvwxyz")

var UpperCaseLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var Numbers = []rune("0123456789")

func RandomString(n int, characterSet []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterSet[rand.Intn(len(characterSet))]
	}
	return string(b)
}

// RandIntRange inclusive
func RandIntRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
