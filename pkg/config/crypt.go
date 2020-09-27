package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// EncryptPayload is a utility function which encrypts a payload using
// AES and a given key.
func EncryptPayload(payload []byte, key []byte) ([]byte, error) {
	// Create an AES instance
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	n, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	if n != len(nonce) {
		return nil, errors.New("Unable to fill nonce")
	}

	return gcm.Seal(nonce, nonce, payload, nil), nil
}

// DecryptPayload is a utility function which decrypts a payload using
// AES and a given key.
func DecryptPayload(ciphertext []byte, key []byte) ([]byte, error) {
	// Create an AES instance
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Use gcm mode (Galois Counter Mode)
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
