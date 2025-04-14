package cryptocloud

type Currency string

const (
	USD Currency = "USD"
	UZS Currency = "UZS"
	KGS Currency = "KGS"
	KZT Currency = "KZT"
	AMD Currency = "AMD"
	AZN Currency = "AZN"
	BYN Currency = "BYN"
	AUD Currency = "AUD"
	TRY Currency = "TRY"
	AED Currency = "AED"
	CAD Currency = "CAD"
	CNY Currency = "CNY"
	HKD Currency = "HKD"
	IDR Currency = "IDR"
	INR Currency = "INR"
	JPY Currency = "JPY"
	PHP Currency = "PHP"
	SGD Currency = "SGD"
	THB Currency = "THB"
	VND Currency = "VND"
	MYR Currency = "MYR"
	RUB Currency = "RUB"
	UAH Currency = "UAH"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

func (c Currency) Ptr() *Currency {
	return &c
}

type CryptoCurrency string

const (
	USDT_TRC20 CryptoCurrency = "USDT_TRC20"
	USDC_TRC20 CryptoCurrency = "USDC_TRC20"
	TUSD_TRC20 CryptoCurrency = "TUSD_TRC20"
	USDT_ERC20 CryptoCurrency = "USDT_ERC20"
	USDC_ERC20 CryptoCurrency = "USDC_ERC20"
	TUSD_ERC20 CryptoCurrency = "TUSD_ERC20"
	BTC        CryptoCurrency = "BTC"
	LTC        CryptoCurrency = "LTC"
	ETH        CryptoCurrency = "ETH"
)

type CryptoCurrencyInfo struct {
	ID              int            `json:"id"`
	Code            string         `json:"code"`
	FullCode        CryptoCurrency `json:"fullcode"`
	Network         Network        `json:"network"`
	Name            string         `json:"name"`
	IsEmailRequired bool           `json:"is_email_required"`
	StableCoin      bool           `json:"stablecoin"`
	IconBase        string         `json:"icon_base"`
	IconNetwork     string         `json:"icon_network"`
	IconQr          string         `json:"icon_qr"`
	Order           int            `json:"order"`
}

type Network struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Icon     string `json:"icon"`
	Fullname string `json:"fullname"`
}
