package tokengen

import (
	"math/rand"
	"time"
)

const tokenLen = 10

func TokenGenerator() string {
	rand.Seed(time.Now().UnixNano())
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	res := make([]rune, tokenLen)
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}
