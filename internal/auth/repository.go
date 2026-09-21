package auth

import (
	"context"
	"errors"
	"time"
)

var ErrSessionNotFound = errors.New("session not found")

type Session struct {
	ID        string
	UserID    int64
	ExpiresAt time.Time
	CreatedAt time.Time
}

type SessionRespository interface {
	Create(ctx context.Context, session *Session) error
	Find(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}
