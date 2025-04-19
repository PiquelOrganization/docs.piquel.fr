package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func VerifySignature(payload, header, secret string) (bool, error) {
	signature, err := hex.DecodeString(header)
	if err != nil {
		return false, err
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))

	expectedSignature := h.Sum(nil)

    log.Printf("Signature verification request -> incoming signatre: %s, expectedSignature: %s", string(signature), string(expectedSignature))
	return hmac.Equal(signature, expectedSignature), nil
}
