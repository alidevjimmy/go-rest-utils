package crypto

import (
	"crypto/sha256"
	"fmt"
)

func GenerateSha256(s string) string {
	return fmt.Sprintf("%x" , sha256.Sum256([]byte(s)))
}

func ValidateSh256(s1 , s2 string) bool {
	p := fmt.Sprintf("%x" , sha256.Sum256([]byte(s1)))
	if p != s2 {
		return false
	}
	return true
}