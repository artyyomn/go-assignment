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
	UpdateLastLogin(ctx context.Context, userID int64, lastLoginAt time.Time) error
}
