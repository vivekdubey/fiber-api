package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghikjklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomBookTitle() string {
	return RandomString(10)
}

func RandomBookAuthor() string {
	return RandomString(10)
}

func RandomDescription() string {
	return RandomString(20)
}

func RandomFirstName() string {
	return RandomString(20)
}
