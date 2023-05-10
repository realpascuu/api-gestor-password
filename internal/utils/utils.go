package utils

import "math/rand"

func GenerateRandomString(n int) (string, error) {
	result := ""
	for i := 0; i <= n; i++ {
		n := rand.Intn(126+1-33) + 33
		result += string(rune(n))
	}
	return result, nil
}
