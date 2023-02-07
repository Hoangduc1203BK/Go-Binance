package model

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Id          uint `gorm:"primaryKey"`
	TokenString string
	Expires     int64
	Blacklisted bool
	UserId uint
}