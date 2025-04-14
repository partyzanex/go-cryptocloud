package cryptocloud

type Statistics struct {
	Count  Statistic[int]     `json:"count"`
	Amount Statistic[float64] `json:"amount"`
}

type Number interface {
	int | float64
}

type Statistic[N Number] struct {
	All      N `json:"all"`
	Created  N `json:"created"`
	Paid     N `json:"paid"`
	Overpaid N `json:"overpaid"`
	Partial  N `json:"partial"`
	Canceled N `json:"canceled"`
}
