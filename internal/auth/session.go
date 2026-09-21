package auth

import (
	"errors"
	"time"
)

var (
	ErrUserExists         = errors.New("username already exists")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrAccountLocked      = errors.New("account is locked")
)

type Config struct {
	SessionTimeout time.Duration
}
