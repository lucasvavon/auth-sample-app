package pg

import (
	"gorm.io/gorm"
	"remember-me/internal/domain/models"
)

type PostgresSessionRepository struct {
	db *gorm.DB
}

func NewPostgresSessionRepository(db *gorm.DB) *PostgresSessionRepository {
	return &PostgresSessionRepository{db}
}

func (r *PostgresSessionRepository) CreateSession(session *models.Session) error {
	req := r.db.Create(&session)
	if req.Error != nil {
		return req.Error
	}
	return nil
}
