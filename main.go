package main

import (
	"binance/database"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnv()
}

func main() {
	database.ConnectionDatabase()
	router := gin.Default()

	router.Run(":" + os.Getenv("PORT"))
}
