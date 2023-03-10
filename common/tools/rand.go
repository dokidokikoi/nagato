package tools

import (
	"math/rand"
	"time"
)

func GetRandStr(n int) string {
	rand.Seed(time.Now().UnixNano())

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < n; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}

	return result
}
