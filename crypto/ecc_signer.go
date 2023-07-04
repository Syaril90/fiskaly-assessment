package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type ECCSigner struct {
	PrivateKey crypto.PrivateKey
}

func NewECCSigner(privateKey crypto.PrivateKey) *ECCSigner {
	return &ECCSigner{
		PrivateKey: privateKey,
	}
}

func (s *ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256([]byte(dataToBeSigned))

	privateKey, ok := s.PrivateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not of type ecdsa.PrivateKey")
	}

	r, sInt, err := ecdsa.Sign(rand.Reader, privateKey, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	signature := append(r.Bytes(), sInt.Bytes()...)

	return signature, nil
}
