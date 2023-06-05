package responsedto

// ErrorResponseData -
type ErrorResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse -
type ErrorResponse struct {
	Success bool              `json:"success"`
	Error   ErrorResponseData `json:"data"`
}

type Response struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type ResponseV2 struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	List    interface{} `json:"list,omitempty"`
}
type Data struct {
	Message string `json:"message"`
}

type FyersAccessTokenResponse struct {
	S           string `json:"s"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
	Id          string `json:"id"`
}

type FyersAuthCodeResponse struct {
	S                  string `json:"s" form:"s"`
	Code               int    `json:"code" form:"code"`
	Authorization_code string `json:"authorization_code" form:"auth_code"`
}

type FyersLiveData struct {
	S string `json:"s"`
	D []struct {
		N string `json:"n"`
		S string `json:"s"`
		V struct {
			Ch             float64 `json:"ch"`
			Chp            float64 `json:"chp"`
			Lp             float64 `json:"lp"`
			Spread         float64 `json:"spread"`
			Ask            float64 `json:"ask"`
			Bid            float64 `json:"bid"`
			OpenPrice      float64 `json:"open_price"`
			HighPrice      float64 `json:"high_price"`
			LowPrice       float64 `json:"low_price"`
			PrevClosePrice float64 `json:"prev_close_price"`
			Volume         int     `json:"volume"`
			ShortName      string  `json:"short_name"`
			Exchange       string  `json:"exchange"`
			Description    string  `json:"description"`
			OriginalName   string  `json:"original_name"`
			Symbol         string  `json:"symbol"`
			FyToken        string  `json:"fyToken"`
			Tt             int     `json:"tt"`
		} `json:"v"`
	} `json:"d"`
}

type FyersOrderInfo struct {
	S         string `json:"s"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	OrderBook []struct {
		OrderDateTime     string  `json:"orderDateTime"`
		ID                string  `json:"id"`
		ExchOrdID         string  `json:"exchOrdId"`
		Side              int     `json:"side"`
		Segment           int     `json:"segment"`
		Instrument        int     `json:"instrument"`
		ProductType       string  `json:"productType"`
		Status            int     `json:"status"`
		Qty               int     `json:"qty"`
		RemainingQuantity int     `json:"remainingQuantity"`
		FilledQty         int     `json:"filledQty"`
		LimitPrice        float64 `json:"limitPrice"`
		StopPrice         float64 `json:"stopPrice"`
		Type              int     `json:"type"`
		DiscloseQty       int     `json:"discloseQty"`
		DqQtyRem          int     `json:"dqQtyRem"`
		OrderValidity     string  `json:"orderValidity"`
		Source            string  `json:"Source"`
		Exchange          int     `json:"exchange"`
		SlNo              int     `json:"slNo"`
		FyToken           string  `json:"fyToken"`
		OfflineOrder      bool    `json:"offlineOrder"`
		Message           string  `json:"message"`
		OrderNumStatus    string  `json:"orderNumStatus"`
		TradedPrice       float64 `json:"tradedPrice"`
		Symbol            string  `json:"symbol"`
		Ch                float64 `json:"ch"`
		Chp               float64 `json:"chp"`
		Lp                float64 `json:"lp"`
		ExSym             string  `json:"ex_sym"`
		Description       string  `json:"description"`
	} `json:"orderBook"`
}
