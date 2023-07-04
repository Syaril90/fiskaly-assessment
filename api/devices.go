package api

import (
	"log"
	"net/http"

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
	LastSignature    string    `json:"last_signature"`
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

func RemapDeviceToResponse(device domain.Device) DeviceResponse {
	return DeviceResponse{
		ID:               device.ID,
		Label:            device.Label,
		SignatureCounter: device.SignatureCounter,
		Algorithm:        device.Algorithm,
		LastSignature:    device.LastSignature,
	}
}

func (s *Server) logError(err error) {
	log.Println(err)
}
