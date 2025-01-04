package mock

import (
	"errors"
	"remember-me/internal/domain/models"
)

type MockUserRepository struct {
	Users       []models.User
	UserByID    *models.User
	UserByEmail *models.User
	SaveError   error
	UpdateError error
	DeleteError error
	ExistsByE   bool
}

// GetUsers retourne tous les utilisateurs
func (m *MockUserRepository) GetUsers() (models.Users, error) {
	return m.Users, nil
}

// GetUserByID retourne un utilisateur par son ID
func (m *MockUserRepository) GetUserByID(id int) (models.User, error) {
	for _, user := range m.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// GetUserByEmail retourne un utilisateur par son email
func (m *MockUserRepository) GetUserByEmail(email string) (models.User, error) {
	for _, user := range m.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

// CreateUser ajoute un utilisateur dans le slice
func (m *MockUserRepository) CreateUser(user *models.User) error {
	if m.SaveError != nil {
		return m.SaveError
	}
	m.Users = append(m.Users, *user)
	return nil
}

// UpdateUser met à jour un utilisateur existant
func (m *MockUserRepository) UpdateUser(id int, user *models.User) error {
	if m.UpdateError != nil {
		return m.UpdateError
	}
	for i, u := range m.Users {
		if u.ID == id {
			m.Users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUser supprime un utilisateur par son ID
func (m *MockUserRepository) DeleteUser(id int) error {
	if m.DeleteError != nil {
		return m.DeleteError
	}
	for i, user := range m.Users {
		if user.ID == id {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// ExistsByEmail vérifie si un utilisateur existe par email
func (m *MockUserRepository) ExistsByEmail(email string) bool {
	return m.ExistsByE
}
