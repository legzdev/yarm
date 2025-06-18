package yarm

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomSuffix() string {
	randomID := make([]byte, 6)

	// rand.Read never returns an error
	rand.Read(randomID)

	return hex.EncodeToString(randomID)
}
