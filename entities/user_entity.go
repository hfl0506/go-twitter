package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsActivate bool   `json:"isActive" gorm:"default: false"`
	ResetCode  string `json:"resetCode" gorm:"default: null"`
}
