package crypto_test

import (
	"crypto/rsa"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/stretchr/testify/assert"
)

func TestRSAGenerator_Generate(t *testing.T) {
	generator := crypto.NewRSAGenerator()

	keyPair, err := generator.Generate()

	assert.NoError(t, err)

	_, okPrivate := keyPair.Private.(*rsa.PrivateKey)
	_, okPublic := keyPair.Public.(*rsa.PublicKey)

	assert.True(t, okPrivate, "Private key should be of type *rsa.PrivateKey")
	assert.True(t, okPublic, "Public key should be of type *rsa.PublicKey")
}
