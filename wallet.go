package cryptocloud

import "github.com/google/uuid"

type Wallet struct {
	Currency *CryptoCurrencyInfo `json:"currency"`
	Address  string              `json:"address"`
	UUID     uuid.UUID           `json:"uuid"`
}
