package main

import (
	"binance/api/users"
	"binance/database"
	model "binance/model/migrate"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnv()
}

func main() {
	database.ConnectionDatabase()
	model.MigrateModel()
	router := gin.Default()
	apiv1 := router.Group("/api/v1")

	users.InitRouter(apiv1)

	router.Run(":" + os.Getenv("PORT"))
}
