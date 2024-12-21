package models

import (
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

type UserDTO struct {
	gorm.Model
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm-password"`
}

func (user *UserDTO) Validate() error {
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
