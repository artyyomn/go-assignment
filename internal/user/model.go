package user

import "time"

type User struct{
	ID int64
	Username string
	PasswordHash string
	FailedAttempts int
	LockedUntil *time.Time
	CreatedAt time.Time
	LastLoginAt *time.Time
} 
