package crypto

import (
	"fmt"
)

type KeyGenerator interface {
	Generate() (*KeyPair, error)
}

func KeyGeneratorFactory(algorithm string) (KeyGenerator, error) {
	switch algorithm {
	case "RSA":
		return &RSAGenerator{}, nil
	case "ECC":
		return &ECCGenerator{}, nil
	default:
		return nil, fmt.Errorf("invalid algorithm")
	}
}
