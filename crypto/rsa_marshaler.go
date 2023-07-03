package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type RSAMarshaler struct{}

func NewRSAMarshaler() Marshaler {
	return &RSAMarshaler{}
}

func (m *RSAMarshaler) Encode(keyPair KeyPair) ([]byte, []byte, error) {

	rsaPublic, ok := keyPair.Public.(*rsa.PublicKey)
	if !ok {
		return nil, nil, ErrInvalidRsaPrivateKey
	}

	rsaPrivate, ok := keyPair.Private.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, ErrInvalidRsaPrivateKey
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(rsaPrivate)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(rsaPublic)

	encodedPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA_PRIVATE_KEY",
		Bytes: privateKeyBytes,
	})

	encodedPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA_PUBLIC_KEY",
		Bytes: publicKeyBytes,
	})

	return encodedPublic, encodedPrivate, nil
}

func (m *RSAMarshaler) Decode(privateKeyBytes []byte) (KeyPair, error) {
	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return KeyPair{}, err
	}

	return KeyPair{
		Private: privateKey,
		Public:  &privateKey.PublicKey,
	}, nil
}
