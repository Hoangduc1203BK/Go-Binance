package model

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	TokenString string
	Expires     int64
	Blacklisted bool
	UserId uint
}