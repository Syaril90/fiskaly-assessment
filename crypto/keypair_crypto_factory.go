package crypto

type KeyGenerator interface {
	Generate() (*KeyPair, error)
}

type Marshaler interface {
	Encode(keyPair KeyPair) ([]byte, []byte, error)
	Decode(privateKeyBytes []byte) (KeyPair, error)
}

func KeyPairCryptoFactory(algorithm string) (KeyGenerator, Marshaler, error) {
	switch algorithm {
	case "RSA":
		return NewRSAGenerator(), NewRSAMarshaler(), nil
	case "ECC":
		return NewECCGenerator(), NewECCMarshaler(), nil
	default:
		return nil, nil, InvalidAlgorithType
	}
}
