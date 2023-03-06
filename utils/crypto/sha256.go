package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 ...
func HashSHA256(data string) string {
	s := sha256.Sum256([]byte(data))
	return hex.EncodeToString(s[:])
}
