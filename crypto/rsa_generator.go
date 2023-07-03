package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

type RSAGenerator struct{}

func NewRSAGenerator() *RSAGenerator {
	return &RSAGenerator{}
}

func (g *RSAGenerator) Generate() (*KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Public:  key.Public(),
		Private: key,
	}, nil
}
