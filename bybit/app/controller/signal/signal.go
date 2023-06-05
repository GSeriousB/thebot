package signal

import (
	"encoding/json"
	"net/http"

	"tradebot/bybit/app/constants"
	"tradebot/bybit/app/controller"
	"tradebot/bybit/app/service/bybit"
	"tradebot/bybit/app/service/logger"

	"tradebot/bybit/app/service/correlation"

	"tradebot/bybit/app/service/dto/request"

	"github.com/gin-gonic/gin"
)

type IWebhookController interface {
	WebhookRequest(c *gin.Context)
}

type WebhookController struct {
	ByBitClient bybit.IByBitConnect
}

func NewWebhookController(byBitClient bybit.IByBitConnect,
) IWebhookController {
	return &WebhookController{
		ByBitClient: byBitClient,
	}
}

func (u WebhookController) WebhookRequest(c *gin.Context) {
	ctx := correlation.WithReqContext(c)
	log := logger.Logger(ctx)

	dataFromBody := request.WebhookRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(&dataFromBody)
	if err != nil {
		log.Errorf("Invalid Request Body", err)
		controller.RespondWithError(c, http.StatusBadRequest, "invalid request body")
		return
	}

	switch dataFromBody.Alert {
	case constants.BULLISH_CONFIRMATION_PLUS.String():
		if err := u.ByBitClient.DoBullishConfirmationPlus(ctx, dataFromBody); err != nil {
			log.Error("error while placing bullish confirmation plus")
			controller.RespondWithError(c, http.StatusInternalServerError, "Error while placing confirmation plus")
			return
		}
	case constants.CONFIRMATION_PLUS_EXITS_BULLISH_EXIT.String():
		if err := u.ByBitClient.DoConfirmationPlusExistBullish(ctx, dataFromBody); err != nil {
			log.Error("error while placing bullish confirmation plus")
			controller.RespondWithError(c, http.StatusInternalServerError, "Error while placing confirmation plus")
			return
		}
	case constants.BEARISH_CONFIRMATION_PLUS.String():
		if err := u.ByBitClient.DoBearishConfirmationPlus(ctx, dataFromBody); err != nil {
			log.Error("error while placing bullish confirmation plus")
			controller.RespondWithError(c, http.StatusInternalServerError, "Error while placing confirmation plus")
			return
		}
	case constants.CONFIRMATION_PLUS_EXITS_BEARISH_EXIT.String():
		if err := u.ByBitClient.DoConfirmationPlusExistBearish(ctx, dataFromBody); err != nil {
			log.Error("error while placing bullish confirmation plus")
			controller.RespondWithError(c, http.StatusInternalServerError, "Error while placing confirmation plus")
			return
		}
	}

	controller.RespondWithSuccess(c, http.StatusAccepted, "Signal Received", nil)

}
