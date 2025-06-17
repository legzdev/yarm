package yarm

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomID() string {
	randomID := make([]byte, 6)

	_, err := rand.Read(randomID)
	if err == nil {
		return hex.EncodeToString(randomID)
	}

	// TODO: fallback (paranoia)
	return ""
}
