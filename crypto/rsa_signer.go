package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type RSASigner struct {
	PrivateKey crypto.PrivateKey
}

func NewRSASigner(privateKey crypto.PrivateKey) *RSASigner {
	return &RSASigner{
		PrivateKey: privateKey,
	}
}

func (s *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	hashed := sha256.Sum256([]byte(dataToBeSigned))

	privateKey, ok := s.PrivateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("faild to convert the private key")
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign data: %w", err)
	}

	return signature, nil
}
