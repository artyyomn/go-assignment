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
	ErrTOTPRequired       = errors.New("two-factor authentication code is required")
	ErrInvalidTOTP        = errors.New("invalid two-factor authentication code")
	ErrTOTPAlreadyEnabled = errors.New("two-factor authentication is already enabled")
)

type Config struct {
	SessionTimeout time.Duration
}
