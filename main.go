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
	router.GET("/users/:id", users.ListUserByID)
	router.PATCH("/users/:id", users.UpdateUserByID)
	router.DELETE("/users/:id", users.DeleteUserByID)

	router.Run(":" + os.Getenv("PORT"))
}
