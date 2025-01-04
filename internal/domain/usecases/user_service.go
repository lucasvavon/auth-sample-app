package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"remember-me/internal/domain/models"
	"remember-me/internal/domain/ports"
	"remember-me/internal/utils"
)

//go:generate mockgen -source=user_service.go -destination=mock/user_service.go

type UserService struct {
	ur ports.UserRepository
}

func NewUserService(ur ports.UserRepository) *UserService {
	return &UserService{ur: ur}
}

func (s *UserService) GetUsers() (*models.Users, error) {
	return s.ur.GetUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
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

func (s *UserService) UpdateUser(id int, user *models.User) error {
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
