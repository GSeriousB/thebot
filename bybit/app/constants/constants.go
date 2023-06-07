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

type (
	Environment     string
	CONTEXT_KEY     string
	CORRELATION_KEY string
	TRADE_ACTION    string
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
