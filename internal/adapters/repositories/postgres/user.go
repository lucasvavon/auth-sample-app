package postgres

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"remember-me/internal/domain/models"
)

type UserGORMRepository struct {
	db *gorm.DB
}

func NewUserGORMRepository(db *gorm.DB) *UserGORMRepository {
	return &UserGORMRepository{db: db}
}

func (r *UserGORMRepository) GetUsers() (models.Users, error) {
	var users models.Users

	req := r.db.Find(&users)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("messages not found: %v", req.Error))
	}

	return users, nil
}

func (r *UserGORMRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	req := r.db.First(&user, id)
	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of models.User.
		return models.User{}, fmt.Errorf("user not found: %v", req.Error)
	}

	return user, nil
}

func (r *UserGORMRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	req := r.db.First(&user, email)
	if req.Error != nil {
		// Use fmt.Errorf for error formatting and return the zero value of models.User.
		return models.User{}, fmt.Errorf("user not found: %v", req.Error)
	}

	return user, nil
}

func (r *UserGORMRepository) CreateUser(user *models.User) error {
	req := r.db.Create(&user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *UserGORMRepository) DeleteUser(id int) error {
	var user models.User

	req := r.db.Unscoped().Delete(&user, &id)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *UserGORMRepository) UpdateUser(id int, user *models.User) error {
	user.ID = id

	req := r.db.Save(user)

	if req.Error != nil {
		return req.Error
	}

	return nil
}

func (r *UserGORMRepository) ExistsByEmail(email string) (bool, error) {
	_, err := r.GetUserByEmail(email)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}

		return false, nil
	}

	return true, nil
}
