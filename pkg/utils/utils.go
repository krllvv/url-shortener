package utils

import (
	"math/rand"
	"time"
)

func GetRandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	symbols := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	ans := make([]rune, length)
	for i := range ans {
		ans[i] = symbols[r.Intn(len(symbols))]
	}
	return string(ans)
}
