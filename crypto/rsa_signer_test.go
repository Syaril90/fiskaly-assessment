package crypto_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/stretchr/testify/assert"
)

func TestRSASigner_Sign_ValidInput(t *testing.T) {
	testKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer := crypto.NewRSASigner(testKey)
	data := []byte("Test Data")

	signature, err := signer.Sign(data)

	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestRSASigner_Sign_NilInput(t *testing.T) {
	testKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer := crypto.NewRSASigner(testKey)

	signature, err := signer.Sign(nil)

	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestRSASigner_Sign_EmptyInput(t *testing.T) {
	testKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	signer := crypto.NewRSASigner(testKey)

	signature, err := signer.Sign([]byte(""))

	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestRSASigner_Sign_NonRSAPrivateKey(t *testing.T) {
	nonRSAKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	signer := crypto.NewRSASigner(nonRSAKey)
	data := []byte("Test Data")

	_, err := signer.Sign(data)

	assert.Error(t, err)
}
