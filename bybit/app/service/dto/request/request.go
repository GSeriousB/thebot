package request

type WebhookRequest struct {
	Alert  string `json:"alert"`
	Ticker string `json:"ticker"`
	Tf     string `json:"tf"`
	Ohlcv  struct {
		Open   float64 `json:"open"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Close  float64 `json:"close"`
		Volume int     `json:"volume"`
	} `json:"ohlcv"`
	Bartime int64 `json:"bartime"`
}
