package crypto

import (
	"crypto"
)

type KeyPair struct {
	Public  crypto.PublicKey
	Private crypto.PrivateKey
}
