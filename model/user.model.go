package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id          uint `gorm:"primaryKey"`
	Email       string
	Name        string
	Password    string
	PhoneNumber string
}
