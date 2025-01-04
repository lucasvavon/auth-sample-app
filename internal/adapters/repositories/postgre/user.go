package postgre

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"remember-me/internal/domain/models"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) GetUsers() (models.Users, error) {
	var users models.Users

	req := r.db.Find(&users)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}

	return users, nil
}

func (r *GormUserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	req := r.db.First(&user, id)
	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of models.User.
		return models.User{}, fmt.Errorf("user not found: %v", req.Error)
	}

	return user, nil
}

func (r *GormUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	req := r.db.First(&user, "email = ?", email)

	if req.Error != nil {
		if errors.Is(req.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %v", req.Error)
	}
	return &user, nil
}

func (r *GormUserRepository) CreateUser(user *models.User) error {
	req := r.db.Create(&user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *GormUserRepository) DeleteUser(id int) error {
	var user models.User

	req := r.db.Unscoped().Delete(&user, &id)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *GormUserRepository) UpdateUser(id int, user *models.User) error {
	user.ID = id

	req := r.db.Save(user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}
