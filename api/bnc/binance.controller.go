package bnc

import (
	"github.com/gin-gonic/gin"
)

func getTokenPrice(c *gin.Context) {
	query := c.Query("symbol")

	result, err := ServiceGetTokenPrice(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func listTokenPrice(c *gin.Context) {
	var query ListTokenPriceDTO
	query.trend = c.Query("trend")
	query.percent = c.Query("percent")
	query.time = c.Query("time")

	result, err := ServiceListTokenPrice(&query)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func getKlines(c *gin.Context) {
	symbol := "ETHBTC"
	period := "15m"
	result, _ := ServiceGetKline(&symbol, &period)

	c.JSON(200, gin.H{
		"result": result,
	})
}
