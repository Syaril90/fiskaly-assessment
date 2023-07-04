package persistence

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
)

type Repository struct {
	Devices      sync.Map
	Transactions sync.Map
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) SaveDevice(d domain.Device) error {
	if _, loaded := r.Devices.LoadOrStore(d.ID.String(), d); loaded {
		log.Printf("Device with ID: %v already exists, overwriting with new data", d.ID.String())
	}
	return nil
}

func (r *Repository) GetAllDevices() ([]domain.Device, error) {
	var devices []domain.Device

	r.Devices.Range(func(key, value interface{}) bool {
		device, ok := value.(domain.Device)
		if !ok {
			return false
		}

		devices = append(devices, device)

		return true
	})

	if len(devices) == 0 {
		return nil, fmt.Errorf("no devices found")
	}

	return devices, nil
}

func (r *Repository) GetDevice(id uuid.UUID) (domain.Device, error) {
	v, ok := r.Devices.Load(id.String())
	if !ok {
		log.Printf("Device with ID: %v not found", id.String())
		return domain.Device{}, ErrorDeviceNotFound
	}
	return v.(domain.Device), nil
}

func (r *Repository) SaveTransaction(t domain.Transaction) error {
	key := t.DeviceID.String() + t.Data

	if _, loaded := r.Transactions.LoadOrStore(key, t); loaded {
		log.Printf("Transaction with Key: %v already exists, overwriting with new data", key)
	}

	return nil
}

func (r *Repository) UpdateLastSignatureAndCounter(deviceID uuid.UUID, lastSignature string) error {
	value, ok := r.Devices.Load(deviceID.String())
	if !ok {
		return fmt.Errorf("device with ID: %v not found", deviceID.String())
	}

	device, ok := value.(domain.Device)
	if !ok {
		return fmt.Errorf("failed to assert value to Device type")
	}

	device.LastSignature = lastSignature
	device.SignatureCounter++

	r.Devices.Store(deviceID.String(), device)

	return nil
}

func (r *Repository) GetTransactions(id uuid.UUID) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	r.Transactions.Range(func(k, v interface{}) bool {
		t := v.(domain.Transaction)
		if t.DeviceID == id {
			transactions = append(transactions, t)
		}
		return true
	})

	if len(transactions) == 0 {
		log.Printf("No transactions found for device with ID: %v", id.String())
		return nil, ErrorTransactionsNotFound
	}

	return transactions, nil
}
