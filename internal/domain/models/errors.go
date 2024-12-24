package models

import "errors"

// Entity Errors
var (
	ErrEmptyUserField     = errors.New("email and password can't be empty")
	ErrFieldWithSpaces    = errors.New("email and password can't have spaces")
	ErrShortPassword      = errors.New("password shorter than 6 characters")
	ErrLongPassword       = errors.New("password longer than 72 characters")
	ErrInvalidEmail       = errors.New("invalid email address")
	ErrPasswordEq         = errors.New("passwords are not equal")
	ErrInvalidCredentials = errors.New("email or password invalid")
)
