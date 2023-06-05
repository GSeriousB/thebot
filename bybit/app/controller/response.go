package controller

import (
	response "tradebot/bybit/app/service/dto/response"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, response.ErrorResponse{Success: false, Error: response.ErrorResponseData{Code: code, Message: message}})
}

func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, response.ResponseV2{Success: true, Message: message, Data: data})
}
