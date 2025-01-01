package services

import (
	"context"
	"remember-me/internal/domain/ports"
	"time"
)

type SessionService struct {
	sessionRepo ports.SessionRepository
}

func NewSessionService(sessionRepo ports.SessionRepository) *SessionService {
	return &SessionService{sessionRepo: sessionRepo}
}

func (s *SessionService) CreateSession(ctx context.Context, sessionID string, userID int) error {
	ttl := 24 * time.Hour
	return s.sessionRepo.SaveSession(ctx, sessionID, userID, ttl)
}

func (s *SessionService) ValidateSession(ctx context.Context, sessionID string) (string, error) {
	return s.sessionRepo.GetSession(ctx, sessionID)
}

func (s *SessionService) InvalidateSession(ctx context.Context, sessionID string) error {
	return s.sessionRepo.DeleteSession(ctx, sessionID)
}
