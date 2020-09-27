package config

import (
	"crypto/rand"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypt(t *testing.T) {
	key := keyGen()

	clearText := []byte("Mary had a little lamb")

	cipherText, err := EncryptPayload(clearText, key)
	assert.Nil(t, err)

	recoveredText, err := DecryptPayload(cipherText, key)
	assert.Nil(t, err)

	assert.Equal(t, clearText, recoveredText)

}

func keyGen() []byte {
	keySize := 256 / 8
	keyBuffer := make([]byte, keySize)

	n, err := rand.Read(keyBuffer)
	if err != nil {
		log.Fatalf("Unable to generate key: %v", err)
	}

	// This shouldn't happen but it doesn't hurt to be careful
	if n < len(keyBuffer) {
		log.Fatalf("Wrote too few bytes: %v", err)
	}

	return keyBuffer
}
