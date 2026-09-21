package auth

import (
	"errors"
	"time"
)

var (
	ErrUserExists         = errors.New("username already exists")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type Config struct {
	SessionTimeout time.Duration
}
