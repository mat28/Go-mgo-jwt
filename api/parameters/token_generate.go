package parameters

import (
	"math/rand"
	"time"
)

const lettersBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes (n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = lettersBytes[rand.Intn(len(lettersBytes))]
	}
	return string(b)
}
