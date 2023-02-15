package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	// dsn := "host=localhost user=postgres password=postgres dbname=go_binance port=5444 sslmode=disable"
	dsn := "host=postgres user=postgres password=postgres dbname=postgresDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("Fail to connect to DB")
	}

	DB = db
}
