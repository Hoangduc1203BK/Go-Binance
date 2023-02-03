package main

import (
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

	router.Run(":" + os.Getenv("PORT"))
}
