package dtos

type CoinResponseDTO struct {
	Coin   string `json:"coin"`
	Logo   string `json:"logo"`
	Ticker string `json:"ticker"`
	Status string `json:"status"`
	Prices struct {
		BRL string `json:"BRL"`
		USD string `json:"USD"`
	} `json:"prices"`
	Error string `json:"error"`
}
