package ports

import (
	"auth-sample-app/internal/domain/models"
)

type UserRepository interface {
	GetUsers() (*models.Users, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}
