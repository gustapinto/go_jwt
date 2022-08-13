package model

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	jwt.RegisteredClaims `gorm:"-:all"`
	Username             string `json:"username" gorm:"unique"`
	Password             string `json:"password"`
	Token                string `json:"token"`
}
