package persistence

import "sync"

type Repository struct {
	Devices      sync.Map
	Transactions sync.Map
}
