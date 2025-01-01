package ports

import (
	"context"
	"time"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, sessionID string, userID int, ttl time.Duration) error
	GetSession(ctx context.Context, sessionID string) (string, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
