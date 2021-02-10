package util

import (
	"crypto/sha256"
	"fmt"
)
func PasswordHashing(mdp string) string {
	h := sha256.New()
	h.Write([]byte(mdp))
	return fmt.Sprintf("%x", h.Sum(nil))
}