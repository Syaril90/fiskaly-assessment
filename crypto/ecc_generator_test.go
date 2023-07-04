package crypto_test

import (
	"crypto/ecdsa"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/stretchr/testify/assert"
)

func TestECCGenerator_Generate(t *testing.T) {
	t.Run("valid ECC keys", func(t *testing.T) {
		generator := crypto.NewECCGenerator()

		keyPair, err := generator.Generate()

		_, okPrivate := keyPair.Private.(*ecdsa.PrivateKey)
		_, okPublic := keyPair.Public.(*ecdsa.PublicKey)

		assert.NoError(t, err)
		assert.NotNil(t, keyPair)
		assert.True(t, okPrivate, "Private key should be of type *ecdsa.PrivateKey")
		assert.True(t, okPublic, "Public key should be of type *ecdsa.PublicKey")
	})
}
