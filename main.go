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

	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.ListUser)

	router.Run(":" + os.Getenv("PORT"))
}
