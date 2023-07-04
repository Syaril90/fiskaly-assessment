package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TransactionResponse struct {
	Signature string `json:"signature"`
	Data      string `json:"data"`
}

func (s *Server) Transactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId, ok := vars["deviceID"]
	if !ok {
		WriteErrorResponse(w, http.StatusBadRequest, []string{"DeviceID parameter is required"})
		return
	}

	id, err := uuid.Parse(deviceId)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid DeviceID"})
		return
	}

	transactions, err := s.repository.GetTransactions(id)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Error retrieving transactions"})
		return
	}

	var transactionResponses []TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, TransactionResponse{
			Signature: transaction.Signature,
			Data:      transaction.Data,
		})
	}

	jsonResponse, err := json.Marshal(transactionResponses)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Error creating response"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
