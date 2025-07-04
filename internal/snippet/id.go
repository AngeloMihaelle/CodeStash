package snippet

import (
	"crypto/rand"
	"encoding/hex"
)

func generateID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
