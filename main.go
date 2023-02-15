package main

import (
	"binance/api/auth"
	"binance/api/users"
	"binance/database"
	model "binance/model/migrate"
	websocket "binance/websocket"
	"os"

	"github.com/gin-gonic/gin"
	"binance/api/bnc"
)

func init() {
	database.LoadEnv()
}

func main() {
	database.ConnectionDatabase()
	model.MigrateModel()
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/ws", func(c *gin.Context) {
		websocket.WsHandler(c.Writer, c.Request)
	})

	apiv1 := router.Group("/api/v1")

	users.InitRouter(apiv1)
	auth.InitRouter(apiv1)
	bnc.InitRouter(apiv1)

	router.Run(":" + os.Getenv("PORT"))
}
