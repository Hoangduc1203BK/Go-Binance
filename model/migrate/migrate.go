package model

import (
	"binance/database"
	"binance/model"
)

func init() {
	database.ConnectionDatabase()
}

func MigrateModel() {
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Token{})
}
