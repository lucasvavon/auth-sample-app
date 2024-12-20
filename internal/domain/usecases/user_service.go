package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"remember-me/internal/adapters/repositories/postgres"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/ports"
)

type UserService struct {
	ur ports.UserRepository
}

func NewUserService(ur *postgres.UserGORMRepository) *UserService {
	return &UserService{ur: ur}
}

func (s *UserService) GetUsers() (models.Users, error) {
	return s.ur.GetUsers()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.ur.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	return s.ur.GetUserByEmail(email)
}

func (s *UserService) CreateUser(user *models.User) error {

	err := user.Validate()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.ur.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user *models.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}
	return s.ur.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.ur.DeleteUser(id)
}

func (s *UserService) ExistsByEmail(email string) (bool, error) {
	return s.ur.ExistsByEmail(email)
}
