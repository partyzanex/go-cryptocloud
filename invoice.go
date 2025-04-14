package cryptocloud

type InvoiceData struct {
	ShopID    string            `json:"shop_id"`
	Amount    float64           `json:"amount"`
	Currency  *Currency         `json:"currency,omitempty"`
	OrderID   *string           `json:"order_id,omitempty"`
	Email     *string           `json:"email,omitempty"`
	AddFields *InvoiceAddFields `json:"add_fields,omitempty"`
}

type TimeToPay struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type InvoiceAddFields struct {
	TimeToPay           *TimeToPay       `json:"time_to_pay,omitempty"`
	EmailToSend         *string          `json:"email_to_send,omitempty"`
	AvailableCurrencies []CryptoCurrency `json:"available_currencies,omitempty"`
	CryptoCurrency      CryptoCurrency   `json:"cryptocurrency,omitempty"`
	Period              Period           `json:"period,omitempty"`
}

type Period string

const (
	PeriodDay   Period = "day"
	PeriodMonth Period = "month"
	PeriodYear  Period = "year"
)

type InvoiceResult struct {
	UUID             string              `json:"uuid"`
	Created          DateTime            `json:"created"`
	Address          string              `json:"address"`
	ExpiryDate       DateTime            `json:"expiry_date"`
	SideCommission   SideCommission      `json:"side_commission"`
	SideCommissionCC SideCommission      `json:"side_commission_cc"`
	Amount           float64             `json:"amount"`
	AmountUSD        float64             `json:"amount_usd"`
	AmountInFiat     float64             `json:"amount_in_fiat"`
	Fee              float64             `json:"fee"`
	FeeUSD           float64             `json:"fee_usd"`
	ServiceFee       float64             `json:"service_fee"`
	ServiceFeeUSD    float64             `json:"service_fee_usd"`
	FiatCurrency     Currency            `json:"fiat_currency"`
	Status           InvoiceStatus       `json:"status"`
	IsEmailRequired  bool                `json:"is_email_required"`
	Link             string              `json:"link"`
	Currency         *CryptoCurrencyInfo `json:"currency"`
	Project          *Project            `json:"project"`
	TestMode         bool                `json:"test_mode"`
}

type SideCommission string

const (
	SideCommissionClient   SideCommission = "client"
	SideCommissionMerchant SideCommission = "merchant"
)

type Project struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Fail    string `json:"fail"`
	Success string `json:"success"`
	Logo    string `json:"logo"`
}

type InvoiceStatus string

const (
	StatusCreated  InvoiceStatus = "created"
	StatusPaid     InvoiceStatus = "paid"
	StatusPartial  InvoiceStatus = "partial"
	StatusOverpaid InvoiceStatus = "overpaid"
	StatusCanceled InvoiceStatus = "canceled"
)

type Invoice struct {
	InvoiceResult

	Received       float64 `json:"received"`
	ReceivedUSD    float64 `json:"received_usd"`
	ToSurcharge    float64 `json:"to_surcharge"`
	ToSurchargeUSD float64 `json:"to_surcharge_usd"`
}
