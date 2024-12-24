package models

import "time"

var Sessions map[string]Session

type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
}

func (s Session) isExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
