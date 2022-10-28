package str

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// RandomString returns [a-z A-Z 0-9] random string
func RandomString(length int) string {
	buffer := make([]byte, length)
	for i := range buffer {
		pos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		buffer[i] = letters[pos.Int64()]
	}
	return string(buffer)
}

// RandomBytes returns random bytes with specified length.
func RandomBytes(length int) ([]byte, error) {
	buffer := make([]byte, length)
	if _, err := rand.Read(buffer); err != nil {
		return nil, err
	}
	return buffer, nil
}

// RandomBase64 returns a URL-safe, base64 encoded
func RandomBase64(length int) (string, error) {
	randomBytes, err := RandomBytes(length)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes), nil
}
