package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

// RSAGenerator generates a RSA key pair.
type RSAGenerator struct{}

// Generate generates a new RSAKeyPair.
func (g *RSAGenerator) Generate() (*KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Public:  key.PublicKey,
		Private: key,
	}, nil
}
