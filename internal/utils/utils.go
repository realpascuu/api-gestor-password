package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	result := ""
	for i := 0; i <= n; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(126+1-33))
		if err != nil {
			continue
		}
		result += string(rune(n.Int64() + 33))
	}
	return result, nil
}
