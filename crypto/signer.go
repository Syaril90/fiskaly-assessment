package crypto

import (
	"crypto"
	"fmt"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

// TODO: implement RSA and ECDSA signing ...
type SignerFactory struct{}

func NewSignerFactory() *SignerFactory {
	return &SignerFactory{}
}

func (sf *SignerFactory) CreateSigner(keyType string, privateKey crypto.PrivateKey) (Signer, error) {
	switch keyType {
	case "RSA":
		return NewRSASigner(privateKey), nil
	case "ECC":
		return NewECCSigner(privateKey), nil
	default:
		return nil, fmt.Errorf("invalid keyType: %s", keyType)
	}
}
