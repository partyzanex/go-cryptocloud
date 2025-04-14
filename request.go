package cryptocloud

type cancelInvoiceRequest struct {
	UUID string `json:"uuid"`
}

type GetInvoicesParams struct {
	Start  Date `json:"start"`
	End    Date `json:"end"`
	Offset int  `json:"offset,omitempty"`
	Limit  int  `json:"limit,omitempty"`
}

type getInvoiceInfoRequest = cancelInvoiceRequest

type GetStatisticsRequest struct {
	Start Date `json:"start"`
	End   Date `json:"end"`
}

type CreateStaticWalletRequest struct {
	ShopID   string `json:"shop_id"`
	Currency string `json:"currency"`
	Identity string `json:"identity"`
}
