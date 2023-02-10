package bnc

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup) {
	binance := router.Group("/binance")
	binance.GET("/token/price", getTokenPrice)
	binance.GET("/list/token/price", listTokenPrice)
}
