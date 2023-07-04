package crypto_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/stretchr/testify/require"
)

func TestNewECCSigner(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	signer := crypto.NewECCSigner(privateKey)
	require.NotNil(t, signer)
	require.Equal(t, privateKey, signer.PrivateKey)
}

func TestECCSigner_Sign(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	signer := crypto.NewECCSigner(privateKey)

	data := []byte("data to be signed")
	signedData, err := signer.Sign(data)
	require.NoError(t, err)
	require.NotNil(t, signedData)

	// Check that the signature is valid
	hashed := sha256.Sum256(data)
	r, s := new(big.Int).SetBytes(signedData[:len(signedData)/2]), new(big.Int).SetBytes(signedData[len(signedData)/2:])
	valid := ecdsa.Verify(&privateKey.PublicKey, hashed[:], r, s)
	require.True(t, valid)
}

func TestECCSigner_Sign_InvalidPrivateKey(t *testing.T) {
	privateKey := "invalid private key"
	signer := crypto.NewECCSigner(privateKey)

	data := []byte("data to be signed")
	_, err := signer.Sign(data)
	require.Error(t, err)
	require.Equal(t, "private key is not of type ecdsa.PrivateKey", err.Error())
}
