package api

import (
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

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
		LastSignature:    device.LastSignature,
	}
}
