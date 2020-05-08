package model

type Currency struct {
	R030         int     `json:"r030"`
	TXT          string  `json:"txt"`
	Rate         float32 `json:"rate"`
	CC           string  `json:"cc"`
	Exchangedate string  `json:"exchangedate"`
}
