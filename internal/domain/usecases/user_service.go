package usecases

import (
	"errors"
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

func (s *UserService) CreateUser(user *models.UserDTO) error {

	if err := user.Validate(); err != nil {
		return err
	}

	exists := s.ur.ExistsByEmail(user.Email)

	if exists {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return s.ur.CreateUser(&u)
}

func (s *UserService) UpdateUser(id int, user *models.UserDTO) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u := models.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return s.ur.UpdateUser(id, &u)
}

func (s *UserService) DeleteUser(id int) error {
	return s.ur.DeleteUser(id)
}

func (s *UserService) ExistsByEmail(email string) bool {
	return s.ur.ExistsByEmail(email)
}
