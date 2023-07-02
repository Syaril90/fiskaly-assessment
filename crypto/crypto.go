package crypto

import (
	"crypto"
	"errors"
)

var ErrInvalidKeyPairType = errors.New("invalid key pair type")

type KeyPair struct {
	Public  crypto.PublicKey
	Private crypto.PrivateKey
}

// KeyMarshaler is an interface for encoding and decoding a KeyPair.
type KeyMarshaler interface {
	Marshal(KeyPair) ([]byte, []byte, error)
	Unmarshal([]byte) (KeyPair, error)
}
