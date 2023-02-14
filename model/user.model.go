package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey";"column:id"`
	Email       string `gorm:"unique"`
	Name        string
	Password    string
	PhoneNumber string 
	Tokens      []Token `gorm:"foreignKey:UserId;references:Id"`
}
