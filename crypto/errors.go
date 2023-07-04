package crypto

import "errors"

var ErrInvalidKeyPairType = errors.New("invalid key pair type")
var ErrInvalidEcdsaPublicKey = errors.New("public key is not of type *ecdsa.PublicKey")
var ErrInvalidEcdsaPrivateKey = errors.New("private key is not of type *ecdsa.PrivateKey")
var ErrInvalidRsaPublicKey = errors.New("public key is not of type *rsa.PublicKey")
var ErrInvalidRsaPrivateKey = errors.New("private key is not of type *rsa.PrivateKey")
var ErrInvalidAlgorithType = errors.New("invalid algorithm type")
