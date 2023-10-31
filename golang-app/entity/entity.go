package entity

type ExchangeRate struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	CurrencyCode string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}
