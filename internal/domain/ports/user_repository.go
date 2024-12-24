package ports

import (
	"remember-me/internal/domain/models"
)

type UserRepository interface {
	GetUsers() (models.Users, error)
	GetUserByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
	ExistsByEmail(email string) bool
}
