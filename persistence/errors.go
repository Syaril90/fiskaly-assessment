package persistence

import "errors"

var ErrorDeviceNotFound = errors.New("device not found")

var ErrorTransactionsNotFound = errors.New("no transactions found for this device")
