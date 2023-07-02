package crypto

import (
	"fmt"
)

type KeyGenerator interface {
	Generate() (*KeyPair, error)
}

type Marshaler interface {
	Marshal(keyPair interface{}) ([]byte, []byte, error)
	Unmarshal(privateKeyBytes []byte) (interface{}, error)
}

func KeyGeneratorFactory(algorithm string) (KeyGenerator, error) {
	switch algorithm {
	case "RSA":
		return NewRSAGenerator(), nil
	case "ECC":
		return NewECCGenerator(), nil
	default:
		return nil, fmt.Errorf("invalid algorithm")
	}
}
