package domain

import (
	"crypto"

	"github.com/google/uuid"
)

// TODO: signature device domain model ...
type Device struct {
	ID               uuid.UUID         `json:"id"`
	Label            string            `json:"label"`
	SignatureCounter int               `json:"signature_counter"`
	LastSignature    string            `json:"last_signature"`
	PrivateKey       crypto.PrivateKey `json:"privateKey"`
	PublicKey        crypto.PublicKey  `json:"publicKey"`
	Algorithm        string            `json:"algorithm"`
}

type Transaction struct {
	DeviceID  uuid.UUID `json:"device_id"`
	Data      string    `json:"data"`
	Signature string    `json:"signature"`
}
