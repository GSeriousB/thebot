package server

import (
	"context"
	"strings"
	"tradebot/bybit/app/apiclient"
	"tradebot/bybit/app/constants"
	"tradebot/bybit/app/service/bybit"
	"tradebot/bybit/app/service/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	healthCheckController "tradebot/bybit/app/controller/healthcheck"
	webhookController "tradebot/bybit/app/controller/signal"
)

func Init(ctx context.Context) {
	if strings.EqualFold(constants.Config.Environment, "prod") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := NewRouter(ctx)

	log := logger.Logger(ctx)
	err := r.Run(constants.Config.HTTPServerConfig.HTTPSERVER_LISTEN + ":" + constants.Config.HTTPServerConfig.HTTPSERVER_PORT)
	if err != nil {
		log.Fatalf("Server not able to startup with error: %v", err)
	}
}

func NewRouter(ctx context.Context) *gin.Engine {
	log := logger.Logger(ctx)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(uuidInjectionMiddleware())

	api := apiclient.NewApiClient()
	byBitClient := bybit.NewByBitConnect(ctx, api)

	healthCheckController := healthCheckController.NewHealthCheckController()
	webhookController := webhookController.NewWebhookController(byBitClient)

	v1 := router.Group("/v1")
	{
		v1.GET(HEALTHCHECK, healthCheckController.HealthCheck)
		v1.POST(WEBHOOK, webhookController.WebhookRequest)
	}
	log.Info("Routes Setup")

	return router
}

// uuidInjectionMiddleware injects the request context with a correlation id of type uuid
func uuidInjectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.GetHeader(constants.CORRELATION_KEY_ID.String())
		if len(correlationId) == 0 {
			correlationID, _ := uuid.NewUUID()
			correlationId = correlationID.String()
			c.Request.Header.Set(constants.CORRELATION_KEY_ID.String(), correlationId)
		}
		c.Writer.Header().Set(constants.CORRELATION_KEY_ID.String(), correlationId)

		c.Next()
	}
}
