package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func NewToken() (string, error) {
	b := make([]byte, 32)
	if _, e := rand.Read(b); e != nil {
		return "", e
	}
	return hex.EncodeToString(b), nil
}
func HashToken(t string) string       { h := sha256.Sum256([]byte(t)); return hex.EncodeToString(h[:]) }
func HashPassword(p string) string    { return HashToken("pwd:" + p) }
func VerifyPassword(h, p string) bool { return h == HashPassword(p) }
func RequireRole(got, want string) error {
	if got != want {
		return fmt.Errorf("forbidden role %s", got)
	}
	return nil
}
