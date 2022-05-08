package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// HmacSha1Sign returns a hexadecimal encoded string
func HmacSha1Sign(text, key string) string {
	byteKey := []byte(key)
	h := hmac.New(sha1.New, byteKey)
	h.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(h.Sum((nil)))
}

// HashPassword returns a hexadecimal encoded string
func HashPassword(original string) string {
	hasher := sha256.New()
	hasher.Write([]byte(original))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
