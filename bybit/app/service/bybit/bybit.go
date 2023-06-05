package bybit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"tradebot/bybit/app/apiclient"
	"tradebot/bybit/app/constants"
	"tradebot/bybit/app/service/dto/request"
	"tradebot/bybit/app/service/logger"
)

type IByBitConnect interface {
	DoBullishConfirmationPlus(ctx context.Context, req request.WebhookRequest) error
	DoBearishConfirmationPlus(ctx context.Context, req request.WebhookRequest) error
	DoConfirmationPlusExistBullish(ctx context.Context, req request.WebhookRequest) error
	DoConfirmationPlusExistBearish(ctx context.Context, req request.WebhookRequest) error
}

type ByBit struct {
	ApiClient    apiclient.IApiClient
	API_KEY      string
	API_SIGN_KEY string
	API_ENDPOINT string
}

type ByBitRequest struct {
	Category    string `json:"category"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderType   string `json:"orderType"`
	Qty         string `json:"qty"`
	Price       string `json:"price"`
	TimeInForce string `json:"timeInForce"`
	OrderLinkID string `json:"orderLinkId"`
	IsLeverage  int    `json:"isLeverage"`
	OrderFilter string `json:"orderFilter"`
}

type BalanceResponse struct {
	Result struct {
		BTC struct {
			AvailableBalance string `json:"available_balance"`
		} `json:"BTC"`
	} `json:"result"`
}

type OrderResponse struct {
	Result struct {
		OrderID string `json:"order_id"`
	} `json:"result"`
}

func sign(secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	return hex.EncodeToString(h.Sum(nil))
}

func NewByBitConnect(ctx context.Context, apiClient apiclient.IApiClient) IByBitConnect {
	return &ByBit{
		ApiClient:    apiClient,
		API_KEY:      constants.Config.ByBitConfig.BYBIT_API_KEY,
		API_SIGN_KEY: sign(constants.Config.ByBitConfig.BYBIT_SECRET_KEY),
		API_ENDPOINT: constants.Config.ByBitConfig.BYBIT_ENDPOINT,
	}
}

func (u ByBit) DoBullishConfirmationPlus(ctx context.Context, req request.WebhookRequest) error {
	log := logger.Logger(ctx)

	// Here, 'ticker' is the asset you want to trade.
	ticker := req.Ticker
	balance, err := u.getBalance(ctx)
	if err != nil {
		log.Errorf("Failed to get balance:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Buy",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place buy order:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Sell",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place buy order:", err)
		return err
	}

	log.Info("Successfully placed orders.")
	return nil
}

func (u ByBit) DoBearishConfirmationPlus(ctx context.Context, req request.WebhookRequest) error {
	log := logger.Logger(ctx)

	// Here, 'ticker' is the asset you want to trade.
	ticker := req.Ticker
	balance, err := u.getBalance(ctx)
	if err != nil {
		log.Errorf("Failed to get balance:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Sell",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place sell order:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Buy",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place buy order:", err)
		return err
	}

	log.Info("Successfully placed orders.")
	return nil
}

func (u ByBit) DoConfirmationPlusExistBullish(ctx context.Context, req request.WebhookRequest) error {
	log := logger.Logger(ctx)

	// Here, 'ticker' is the asset you want to trade.
	ticker := req.Ticker
	balance, err := u.getBalance(ctx)
	if err != nil {
		log.Errorf("Failed to get balance:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Sell",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place sell order:", err)
		return err
	}

	log.Info("Successfully placed orders.")
	return nil
}

func (u ByBit) DoConfirmationPlusExistBearish(ctx context.Context, req request.WebhookRequest) error {
	log := logger.Logger(ctx)

	// Here, 'ticker' is the asset you want to trade.
	ticker := req.Ticker
	balance, err := u.getBalance(ctx)
	if err != nil {
		log.Errorf("Failed to get balance:", err)
		return err
	}

	if err := u.placeOrder(ctx, ByBitRequest{
		Category:    "linear",
		Symbol:      ticker,
		IsLeverage:  0,
		Side:        "Buy",
		OrderType:   "Limit",
		Qty:         balance,
		TimeInForce: "GoodTillCancel",
	}); err != nil {
		log.Errorf("Failed to place buy order:", err)
		return err
	}

	log.Info("Successfully placed orders.")
	return nil
}

func (u ByBit) getBalance(ctx context.Context) (string, error) {
	log := logger.Logger(ctx)
	url := fmt.Sprintf("%s/v5/account/wallet-balance?accountType=UNIFIED", u.API_ENDPOINT)

	var headers = make(map[string]string)

	headers["X-BAPI-API-KEY"] = u.API_KEY
	headers["X-BAPI-TIMESTAMP"] = strconv.FormatInt(time.Now().Unix()*1000, 10)
	headers["X-BAPI-RECV-WINDOW"] = "20000"
	headers["X-BAPI-SIGN"] = u.API_SIGN_KEY
	headers["X-BAPI-SIGN-TYPE"] = "2"

	// * Creates a request body
	req, requestError := u.ApiClient.CreateJSONRequest(ctx, http.MethodGet, url, headers, nil)
	if requestError != nil {
		log.Warnw("Error for calling access token err: ", requestError)
		return "", requestError
	}

	// * Makes an api call to the OTP service
	status, response, err := u.ApiClient.RestExecute(ctx, req)
	if err != nil || status != http.StatusOK {
		log.Warnw("Error for calling access token err: ", err, "response :", response)
		return "", err
	}

	var balanceResp BalanceResponse
	err = json.Unmarshal([]byte(response), &balanceResp)
	if err != nil {
		log.Errorf("Error unmarshalling err: ", err)
		return "", err
	}

	return balanceResp.Result.BTC.AvailableBalance, nil
}

func (u ByBit) placeOrder(ctx context.Context, request interface{}) error {
	log := logger.Logger(ctx)
	url := fmt.Sprintf("%s/v5/order/create", u.API_ENDPOINT)

	var headers = make(map[string]string)

	headers["X-BAPI-API-KEY"] = u.API_KEY
	headers["X-BAPI-TIMESTAMP"] = strconv.FormatInt(time.Now().Unix()*1000, 10)
	headers["X-BAPI-RECV-WINDOW"] = "20000"
	headers["X-BAPI-SIGN"] = u.API_SIGN_KEY
	headers["X-BAPI-SIGN-TYPE"] = "2"

	// * Creates a request body
	req, requestError := u.ApiClient.CreateJSONRequest(ctx, http.MethodPost, url, headers, request)
	if requestError != nil {
		log.Warnw("Error for calling access token err: ", requestError)
		return requestError
	}

	// * Makes an api call to the OTP service
	status, response, err := u.ApiClient.RestExecute(ctx, req)
	if err != nil || status != http.StatusOK {
		log.Warnw("Error for calling access token err: ", err, "response :", response)
		return err
	}

	var orderResp OrderResponse
	err = json.Unmarshal([]byte(response), &orderResp)
	if err != nil {
		log.Errorf("Error unmarshalling err: ", err)
		return err
	}

	return nil
}
