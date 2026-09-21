package user

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, user *User) error
	FindUser(ctx context.Context, username string) (*User, error)
	UpdateLoginState(ctx context.Context, userID int64, failedAttempts int, lockedUntil *time.Time) error
	UpdateLastLogin(ctx context.Context, userID int64, lastLoginAt time.Time) error
	UpdateTOTP(ctx context.Context, userID int64, secret *string, enabled bool) error
}
