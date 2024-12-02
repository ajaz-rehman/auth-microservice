package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateRefreshToken() (token string, err error) {
	b := make([]byte, 32)

	_, err = rand.Read(b)

	if err != nil {
		return
	}

	token = hex.EncodeToString(b)

	return
}
