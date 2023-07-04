package persistence_test

import (
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
)

func TestRepository(t *testing.T) {
	repo := persistence.NewRepository()

	tests := []struct {
		name   string
		device domain.Device
		trans  domain.Transaction
	}{
		{
			name: "First Device",
			device: domain.Device{
				ID:               uuid.New(),
				Algorithm:        "algorithm",
				PrivateKey:       []byte("privateKey"),
				SignatureCounter: 0,
			},
			trans: domain.Transaction{
				DeviceID:  uuid.New(),
				Signature: "signature",
				Data:      "data",
			},
		},
		{
			name: "Second Device",
			device: domain.Device{
				ID:               uuid.New(),
				Algorithm:        "algorithm2",
				PrivateKey:       []byte("privateKey2"),
				SignatureCounter: 0,
			},
			trans: domain.Transaction{
				DeviceID:  uuid.New(),
				Signature: "signature2",
				Data:      "data2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// SaveDevice and GetDevice
			err := repo.SaveDevice(test.device)
			assert.NoError(t, err)

			retrievedDevice, err := repo.GetDevice(test.device.ID)
			assert.NoError(t, err)
			assert.Equal(t, test.device, retrievedDevice)

			// GetAllDevices
			devices, err := repo.GetAllDevices()
			assert.NoError(t, err)
			assert.Contains(t, devices, test.device)

			// SaveTransaction and GetTransactions
			err = repo.SaveTransaction(test.trans)
			assert.NoError(t, err)

			transactions, err := repo.GetTransactions(test.trans.DeviceID)
			assert.NoError(t, err)
			assert.Contains(t, transactions, test.trans)

			// UpdateLastSignatureAndCounter
			newSignature := "newSignature"
			err = repo.UpdateLastSignatureAndCounter(test.device.ID, newSignature)
			assert.NoError(t, err)

			updatedDevice, err := repo.GetDevice(test.device.ID)
			assert.NoError(t, err)
			assert.Equal(t, newSignature, updatedDevice.LastSignature)
			assert.Equal(t, test.device.SignatureCounter+1, updatedDevice.SignatureCounter)
		})
	}
}
