package keygen

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateToken for api key
func GenerateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
