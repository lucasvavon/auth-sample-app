package usecases

import (
	"context"
	"remember-me/internal/domain/ports"
	"time"
)

type SessionService struct {
	rs ports.SessionRepository
}

func NewSessionService(rs ports.SessionRepository) *SessionService {
	return &SessionService{rs: rs}
}

func (s *SessionService) CreateSession(ctx context.Context, sessionID string, userID int) error {
	ttl := 24 * time.Hour
	return s.rs.SaveSession(ctx, sessionID, userID, ttl)
}

func (s *SessionService) ValidateSession(ctx context.Context, sessionID string) (string, error) {
	return s.rs.GetSession(ctx, sessionID)
}

func (s *SessionService) InvalidateSession(ctx context.Context, sessionID string) error {
	return s.rs.DeleteSession(ctx, sessionID)
}
