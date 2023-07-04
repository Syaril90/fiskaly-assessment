package api

import "net/http"

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
