package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/artyyomn/go-assignment/internal/user"
)

type Service struct {
	user    user.Repository
	session SessionRespository
	config  Config
}

const MinPasswordLength = 8

const (
	maxLoginAttempts = 3
	lockoutDuration  = 15 * time.Minute
)

func NewService(user user.Repository, session SessionRespository, config Config) *Service {
	return &Service{
		user:    user,
		session: session,
		config:  config,
	}
}

func (s *Service) Register(ctx context.Context, username string, password string) error {
	// validate
	// check existing user
	// hash password
	// construct User
	username = strings.TrimSpace(username)

	if username == "" {
		return ErrInvalidUsername
	}

	if password == "" {
		return ErrInvalidPassword
	}
	if len([]rune(password)) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	// Check whether the username is already registered.
	existingUser, err := s.user.FindUser(ctx, username)
	if err != nil {
		if !errors.Is(err, user.ErrNotFound) {
			return err
		}
	}

	if existingUser != nil {
		return ErrUserExists
	}

	//haishng
	passwordHash, err := HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &user.User{
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}

	return s.user.Create(ctx, newUser)
}

func (s *Service) Login(ctx context.Context, username string, password string) (*user.User, *Session, error) {
	username = strings.TrimSpace(username)

	if username == "" {
		return nil, nil, ErrInvalidUsername
	}

	if password == "" {
		return nil, nil, ErrInvalidPassword
	}

	existingUser, err := s.user.FindUser(ctx, username)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, nil, ErrInvalidCredentials
		}
		return nil, nil, err
	}

	now := time.Now().UTC()
	if existingUser.LockedUntil != nil {
		if now.Before(*existingUser.LockedUntil) {
			return nil, nil, ErrAccountLocked
		}

		if err := s.user.UpdateLoginState(ctx, existingUser.ID, 0, nil); err != nil {
			return nil, nil, err
		}
		existingUser.FailedAttempts = 0
		existingUser.LockedUntil = nil
	}

	if err := VerifyPassword(password, existingUser.PasswordHash); err != nil {
		failedAttempts := existingUser.FailedAttempts + 1
		var lockedUntil *time.Time
		if failedAttempts >= maxLoginAttempts {
			lockExpiry := now.Add(lockoutDuration)
			lockedUntil = &lockExpiry
		}

		if updateErr := s.user.UpdateLoginState(ctx, existingUser.ID, failedAttempts, lockedUntil); updateErr != nil {
			return nil, nil, updateErr
		}
		if lockedUntil != nil {
			return nil, nil, ErrAccountLocked
		}
		return nil, nil, ErrInvalidCredentials
	}
	if s.session == nil {
		return nil, nil, errors.New("session repository is not configured")
	}
	if s.config.SessionTimeout <= 0 {
		return nil, nil, errors.New("session timeout must be greater than zero")
	}

	session := &Session{
		UserID:    existingUser.ID,
		CreatedAt: now,
		ExpiresAt: now.Add(s.config.SessionTimeout),
	}
	if err := s.session.Create(ctx, session); err != nil {
		return nil, nil, err
	}
	if err := s.user.UpdateLoginState(ctx, existingUser.ID, 0, nil); err != nil {
		return nil, nil, err
	}

	lastLoginAt := now
	if err := s.user.UpdateLastLogin(ctx, existingUser.ID, lastLoginAt); err != nil {
		return nil, nil, err
	}
	existingUser.LastLoginAt = &lastLoginAt

	return existingUser, session, nil
}

func (s *Service) Logout(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return ErrSessionNotFound
	}
	if s.session == nil {
		return errors.New("session repository is not configured")
	}

	return s.session.Delete(ctx, sessionID)
}
