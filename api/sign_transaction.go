package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type SignRequest struct {
	Data     string    `json:"data"`
	DeviceId uuid.UUID `json:"deviceId"`
}

type SignResponse struct {
	Signature  string `json:"signature"`
	SignedData string `json:"signed_data"`
}

func (s *Server) Sign(w http.ResponseWriter, r *http.Request) {
	var req SignRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid request body"})
		return
	}

	device, err := s.repository.GetDevice(req.DeviceId)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid request body"})
		return
	}

	privateKeyBytes, ok := device.PrivateKey.([]byte)
	if !ok {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid private key format"})
		return
	}

	keyPairCrypto, err := crypto.NewKeyPairCrypto(device.Algorithm)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Invalid algorithm"})
		return
	}

	keyPair, err := keyPairCrypto.Marshaler.Decode(privateKeyBytes)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Failed to decode private key"})
		return
	}

	signerFactory := crypto.NewSignerFactory()
	signer, err := signerFactory.CreateSigner(device.Algorithm, keyPair.Private)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Failed to decode private key"})
		return
	}

	secureDataToBeSign := createSecureDataToSign(req.Data, device)

	signedData, err := signer.Sign([]byte(secureDataToBeSign))
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusBadRequest, []string{"Failed to sign data"})
		return
	}

	signedDataBase64 := base64.StdEncoding.EncodeToString(signedData)

	err = s.repository.SaveTransaction(domain.Transaction{
		DeviceID:  device.ID,
		Signature: signedDataBase64,
		Data:      fmt.Sprintf("%d_%s_%s", device.SignatureCounter, req.Data, signedDataBase64),
	})
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Failed to save transaction"})
		return
	}

	err = s.repository.UpdateLastSignatureAndCounter(device.ID, signedDataBase64)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Failed to update last signature"})
		return
	}

	res := SignResponse{
		Signature:  signedDataBase64,
		SignedData: secureDataToBeSign,
	}

	device.SignatureCounter++

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Error creating response"})
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		s.logError(err)
		WriteErrorResponse(w, http.StatusInternalServerError, []string{"Error writing response"})
		return
	}
}

func createSecureDataToSign(dataToSign string, device domain.Device) string {
	var lastSignature string

	if device.SignatureCounter == 0 {
		lastSignature = base64.StdEncoding.EncodeToString([]byte(device.ID.String()))
	} else {
		lastSignature = device.LastSignature
	}

	securedDataToBeSigned := fmt.Sprintf("%d_%s_%s", device.SignatureCounter, dataToSign, lastSignature)

	return securedDataToBeSigned
}
