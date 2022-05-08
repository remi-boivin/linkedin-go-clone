package util

import (
	"math/rand"
	"time"
)

// init is used to seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// MakeNonce returns a random string
func MakeNonce(size int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
