package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func VerifySignature(payload, header, secret string) (bool, error) {
	signature, err := hex.DecodeString(header)
	if err != nil {
		return false, err
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))

	exeptectedSignature := h.Sum(nil)

	return hmac.Equal(signature, exeptectedSignature), nil
}
