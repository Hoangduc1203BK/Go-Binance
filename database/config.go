package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "host=localhost user=hmduc password=12345 dbname=gin-api port=5432 sslmode=disable"
	fmt.Print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print("Fail to connect to DB")
	}

	DB = db
}
