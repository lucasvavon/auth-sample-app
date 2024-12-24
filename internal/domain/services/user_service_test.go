package services

import (
	"github.com/stretchr/testify/assert"
	"remember-me/internal/domain/models"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	us := NewUserService(nil)

	err := us.CreateUser(&models.User{
		Email:    "test@example.com",
		Password: "password123",
	})

	assert.NotNil(t, err)

}
