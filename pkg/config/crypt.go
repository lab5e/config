package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// EncryptPayload is a utility function which encrypts a payload using
// AES given a key.
func EncryptPayload(payload []byte, key []byte) ([]byte, error) {
	// Decode the base64 encoded key
	aesKeyDecoded := make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	n, err := base64.StdEncoding.Decode(aesKeyDecoded, key)
	if err != nil {
		return nil, err
	}
	aesKey := aesKeyDecoded[:n]

	// Create an AES instance
	c, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	n, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	if n != len(nonce) {
		return nil, errors.New("Unable to fill nonce")
	}

	return gcm.Seal(nonce, nonce, payload, nil), nil
}
