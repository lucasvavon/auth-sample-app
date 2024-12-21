package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type Users []User
