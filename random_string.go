package utils

import (
	"crypto/rand"
	"math/big"
)

var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var LowercaseLetters = []rune("abcdefghijklmnopqrstuvwxyz")

var UpperCaseLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var Numbers = []rune("0123456789")

func CharacterSet(sets ...[]rune) (result []rune) {
	for _, set := range sets {
		result = append(result, set...)
	}
	return result
}

func RandomString(n int, characterSet []rune) (string, error) {
	b := make([]rune, n)
	for i := range b {
		randIndex, err := RandIntRange(0, len(characterSet)-1)
		if err != nil {

		}
		b[i] = characterSet[randIndex]
	}
	return string(b), nil
}

// RandIntRange inclusive
func RandIntRange(min, max int) (int, error) {
	delta := max - min + 1
	n, err := rand.Int(rand.Reader, big.NewInt(int64(delta)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64() + int64(min)), nil
}
