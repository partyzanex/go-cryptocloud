package cryptocloud

type Balance struct {
	Currency            *CryptoCurrencyInfo `json:"currency"`
	BalanceCrypto       float64             `json:"balance_crypto"`
	BalanceUSD          float64             `json:"balance_usd"`
	AvailableBalance    float64             `json:"available_balance"`
	AvailableBalanceUSD float64             `json:"available_balance_usd"`
}
