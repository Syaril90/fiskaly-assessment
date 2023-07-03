package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type CreateDeviceRequest struct {
	Algorithm string `json:"algorithm"`
	Label     string `json:"label"`
}

type DeviceResponse struct {
	ID               uuid.UUID `json:"id"`
	Label            string    `json:"label"`
	SignatureCounter int       `json:"signature_counter"`
	Algorithm        string    `json:"algorithm"`
}

func (s *Server) Devices(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.CreateDevice(w, r)
	case http.MethodGet:
		s.GetAllDevices(w, r)
	default:
		WriteErrorResponse(w, http.StatusMethodNotAllowed, []string{"Only POST and GET requests are allowed"})
	}
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

	keyPairGenerator, KeyPairMarshaler, err := crypto.KeyPairCryptoFactory(req.Algorithm)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid algorithm"})
		return
	}

	keyPair, err := keyPairGenerator.Generate()
	if err != nil {
		s.logError(err)
		WriteInternalError(w)
		return
	}

	encodePublicKey, encodePrivateKey, err := KeyPairMarshaler.Encode(*keyPair)
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

func (s *Server) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := s.repository.GetAllDevices()
	if err != nil {
		if err.Error() == "no devices found" {
			WriteErrorResponse(w, http.StatusNotFound, []string{"Unable to find any results"})
		} else {
			s.logError(err)
			WriteInternalError(w)
			return
		}
	}

	deviceResponses := make([]DeviceResponse, 0, len(devices))
	for _, device := range devices {
		deviceResponses = append(deviceResponses, RemapDeviceToResponse(device))
	}

	WriteAPIResponse(w, http.StatusOK, deviceResponses)
}

func RemapDeviceToResponse(device domain.Device) DeviceResponse {
	return DeviceResponse{
		ID:               device.ID,
		Label:            device.Label,
		SignatureCounter: device.SignatureCounter,
		Algorithm:        device.Algorithm,
	}
}

func (s *Server) logError(err error) {
	log.Println(err)
}
