package models

import (
	"fmt"
	"gorm.io/gorm"
	"net/mail"
	"strings"
)

type User struct {
	gorm.Model
	ID       int `gorm:"primary_key"`
	Email    string
	Password string
}

type Users []User
type UserDTO struct {
	Email           string `json:"email" form:"email" query:"email"`
	Password        string `json:"password" form:"password" query:"password"`
	ConfirmPassword string `json:"confirm-password" form:"confirm-password" query:"confirm-password"`
}

func (user *User) Validate() error {
	fmt.Printf("\n\ndata : %s - %s \n\n\n", user.Password, user.Email)
	if user.Password == "" || user.Email == "" {
		fmt.Print("\n\nErrEmptyUserField :\n\n")
		return ErrEmptyUserField
	}

	if strings.ContainsAny(user.Password, "\t\r\n") {
		fmt.Println("\n\nErrFieldWithSpaces :\n\n", user.Password)
		return ErrFieldWithSpaces
	}

	if len(user.Password) < 6 {
		fmt.Println("\n\nErrShortPassword :\n\n", user.Password)
		return ErrShortPassword
	}

	if len(user.Password) > 72 {
		fmt.Println("\n\nErrLongPassword :\n\n", user.Password)
		return ErrLongPassword
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		fmt.Println("\n\nErrInvalidEmail :\n\n", user.Password)
		return ErrInvalidEmail
	}

	return nil
}
