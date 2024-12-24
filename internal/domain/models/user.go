package models

import (
	"net/mail"
	"strings"
	"time"
)

type (
	User struct {
		ID              int       `json:"id" gorm:"primaryKey"`
		Email           string    `json:"email" gorm:"unique;not null"`
		Password        string    `json:"password" gorm:"not null"`
		ConfirmPassword string    `json:"confirm-password" gorm:"-"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}

	Users []User
)

func (user *User) Validate() error {
	if user.Password == "" || user.Email == "" {
		return ErrEmptyUserField
	}

	if user.Password != user.ConfirmPassword {
		return ErrPasswordEq
	}

	if strings.ContainsAny(user.Password, "\t\r\n") {
		return ErrFieldWithSpaces
	}

	if len(user.Password) < 6 {
		return ErrShortPassword
	}

	if len(user.Password) > 72 {
		return ErrLongPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
