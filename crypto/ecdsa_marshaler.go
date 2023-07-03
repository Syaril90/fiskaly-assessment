package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

type ECCMarshaler struct{}

func NewECCMarshaler() Marshaler {
	return &ECCMarshaler{}
}

func (m *ECCMarshaler) Encode(keyPair KeyPair) ([]byte, []byte, error) {
	eccPublic, ok := keyPair.Public.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, ErrInvalidEcdsaPublicKey
	}

	eccPrivate, ok := keyPair.Private.(*ecdsa.PrivateKey)
	if !ok {
		return nil, nil, ErrInvalidEcdsaPrivateKey
	}

	privateKeyBytes, err := x509.MarshalECPrivateKey(eccPrivate)
	if err != nil {
		return nil, nil, err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(eccPublic)
	if err != nil {
		return nil, nil, err
	}

	encodedPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE_KEY",
		Bytes: privateKeyBytes,
	})

	encodedPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC_KEY",
		Bytes: publicKeyBytes,
	})

	return encodedPublic, encodedPrivate, nil
}

func (m *ECCMarshaler) Decode(privateKeyBytes []byte) (KeyPair, error) {
	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		Private: privateKey,
		Public:  &privateKey.PublicKey,
	}, nil
}
