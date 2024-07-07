package token

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateToken(byteCount int) string {
	bytes := make([]byte, 15)
	rand.Read(bytes)
	return base32.StdEncoding.EncodeToString(bytes)

}
func DecodeToken(token string) string {
	decoded, err := base32.StdEncoding.DecodeString(token)
	if err != nil {
		return ""
	}
	return string(decoded)
}
