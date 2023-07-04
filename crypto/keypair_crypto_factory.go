package crypto

type KeyGenerator interface {
	Generate() (*KeyPair, error)
}

type Marshaler interface {
	Encode(keyPair KeyPair) ([]byte, []byte, error)
	Decode(privateKeyBytes []byte) (KeyPair, error)
}

type KeyPairCrypto struct {
	KeyGenerator KeyGenerator
	Marshaler    Marshaler
}

func NewKeyPairCrypto(algorithm string) (*KeyPairCrypto, error) {
	switch algorithm {
	case "RSA":
		return &KeyPairCrypto{
			KeyGenerator: NewRSAGenerator(),
			Marshaler:    NewRSAMarshaler(),
		}, nil
	case "ECC":
		return &KeyPairCrypto{
			KeyGenerator: NewECCGenerator(),
			Marshaler:    NewECCMarshaler(),
		}, nil
	default:
		return nil, ErrInvalidAlgorithType
	}
}
