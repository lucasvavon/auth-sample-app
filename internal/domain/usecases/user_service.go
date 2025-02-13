package usecases

import (
	"auth-sample-app/internal/domain/models"
	"auth-sample-app/internal/domain/ports"
	"auth-sample-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur ports.UserRepository
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{ur: ur}
}

func (s *UserService) GetUsers() (*models.Users, error) {
	return s.ur.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.ur.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.ur.GetUserByEmail(email)
}

func (s *UserService) CreateUser(user *models.User) error {

	if err := user.Validate(); err != nil {
		return err
	}

	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return err
	}

	u := models.User{
		Email:    user.Email,
		Password: hashedPassword,
	}

	return s.ur.CreateUser(&u)
}

func (s *UserService) UpdateUser(id uint, user *models.User) error {
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

func (s *UserService) DeleteUser(id uint) error {
	return s.ur.DeleteUser(id)
}
