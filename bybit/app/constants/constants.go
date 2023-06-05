package constants

import (
	"tradebot/bybit/config"
)

const (
	CTK_CLAIM_KEY      = CONTEXT_KEY("claims")
	CORRELATION_KEY_ID = CORRELATION_KEY("X-Correlation-ID")
	//Header constants
	AUTHORIZATION = "Authorization"
	BEARER        = "Bearer "
)

var Config *config.CoreServiceConfig

var AvailableAccountType = []string{}

var DBLOGMODE bool

type (
	Environment     string
	CONTEXT_KEY     string
	CORRELATION_KEY string
	TRADE_ACTION    string
)

var (
	Dev   Environment = "dev"
	Prod  Environment = "prod"
	Stage Environment = "stage"
)

func (c Environment) String() string {
	return string(c)
}

func (c CONTEXT_KEY) String() string {
	return string(c)
}

func (c CORRELATION_KEY) String() string {
	return string(c)
}

var (
	BULLISH_CONFIRMATION_PLUS            TRADE_ACTION = "Bullish Confirmation+"
	CONFIRMATION_PLUS_EXITS_BULLISH_EXIT TRADE_ACTION = "Confirmation + Exits Bullish Exit"
	BEARISH_CONFIRMATION_PLUS            TRADE_ACTION = "Bearish Confirmation+"
	CONFIRMATION_PLUS_EXITS_BEARISH_EXIT TRADE_ACTION = "Confirmation + Exits Bearish Exit"
)

func (c TRADE_ACTION) String() string {
	return string(c)
}

var (
	FYERS_AUTH_CODE          = "https://api.fyers.in/api/v2/generate-authcode?client_id=%s&redirect_uri=%s&response_type=code&state=%s"
	FYERS_VALIDATE_AUTH_CODE = "https://api.fyers.in/api/v2/validate-authcode"
	FYERS_PLACE_ORDER        = "https://api.fyers.in/api/v2/orders"
	FYERS_EXIT_ORDER         = "https://api.fyers.in/api/v2/positions"
	FYERS_ACCESS_TOKEN       = "fyers_access_token"
	FYERS_LIVE_DATA          = "https://api.fyers.in/data-rest/v2/quotes/?symbols=%s"
	FYERS_ORDER_INFO         = "https://api.fyers.in/api/v2/orders?id=%s"
)

// var FYERS_ACCESS_TOKEN_MAPPER = make(map[string]string)
