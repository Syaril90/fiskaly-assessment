package api

import (
	"log"

	"github.com/google/uuid"
)

func (s *Server) logError(err error) {
	log.Println(err)
}

type DeviceResponse struct {
	ID               uuid.UUID `json:"id"`
	Label            string    `json:"label"`
	SignatureCounter int       `json:"signature_counter"`
	Algorithm        string    `json:"algorithm"`
	LastSignature    string    `json:"last_signature"`
}
