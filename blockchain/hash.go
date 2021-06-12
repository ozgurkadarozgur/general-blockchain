package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncodeSha256(data []byte) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
