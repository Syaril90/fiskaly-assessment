package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

func (s *Server) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var req CreateDeviceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid request body"})
		return
	}

	if req.Label == "" {
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Label is required"})
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Unable to create UUID for the device"})
		return
	}

	keyPairCrypto, err := crypto.NewKeyPairCrypto(req.Algorithm)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid algorithm"})
		return
	}

	keyPair, err := keyPairCrypto.KeyGenerator.Generate()
	if err != nil {
		s.logError(err)
		WriteInternalError(w)
		return
	}

	encodePublicKey, encodePrivateKey, err := keyPairCrypto.Marshaler.Encode(*keyPair)
	if err != nil {
		s.logError(err)
		WriteInternalError(w)
		return
	}

	device := domain.Device{
		ID:               id,
		Label:            req.Label,
		SignatureCounter: 0,
		PrivateKey:       encodePrivateKey,
		PublicKey:        encodePublicKey,
		Algorithm:        req.Algorithm,
	}

	err = s.repository.SaveDevice(device)
	if err != nil {
		s.logError(err)
		WriteInternalError(w)
		return
	}

	WriteAPIResponse(w, http.StatusCreated, device)
}
