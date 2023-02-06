package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey"`
	Email       string `gorm:"unique"`
	Name        string
	Password    string
	PhoneNumber string 
}
