package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// ECCGenerator generates an ECC key pair.
type ECCGenerator struct{}

func NewECCGenerator() *ECCGenerator {
	return &ECCGenerator{}
}

// Generate generates a new ECCKeyPair.
func (g *ECCGenerator) Generate() (*KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Public:  key.PublicKey,
		Private: key,
	}, nil
}
