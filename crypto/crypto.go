package crypto

import (
	"crypto"
	"errors"
)

var ErrInvalidKeyPairType = errors.New("invalid key pair type")
var ErrInvalidEcdsaPublicKey = errors.New("public key is not of type *ecdsa.PublicKey")
var ErrInvalidEcdsaPrivateKey = errors.New("private key is not of type *ecdsa.PrivateKey")
var ErrInvalidRsaPublicKey = errors.New("public key is not of type *rsa.PublicKey")
var ErrInvalidRsaPrivateKey = errors.New("private key is not of type *rsa.PrivateKey")
var InvalidAlgorithType = errors.New("invalid algorithm")

type KeyPair struct {
	Public  crypto.PublicKey
	Private crypto.PrivateKey
}

// KeyMarshaler is an interface for encoding and decoding a KeyPair.
type KeyMarshaler interface {
	Marshal(KeyPair) ([]byte, []byte, error)
	Unmarshal([]byte) (KeyPair, error)
}
