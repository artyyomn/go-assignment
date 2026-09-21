package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/artyyomn/go-assignment/internal/user"
)

var (
	ErrUserExists         = errors.New("username already exists")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type Service struct {
	user    user.Repository
	session SessionRespository
}

func NewService(user user.Repository) *Service {
	return &Service{
		user: user,
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

	// Hash the password before storing it.
	passwordHash, err := HashPassword(password)
	if err != nil {
		return err
	}

	newUser := &user.User{
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}

	return s.user.Create(ctx, newUser) // repository.Create()
}

func (s *Service) Login(ctx context.Context, username string, password string) (*user.User, error) {
	username = strings.TrimSpace(username)

	if username == "" {
		return nil, ErrInvalidUsername
	}

	if password == "" {
		return nil, ErrInvalidPassword
	}

	existingUser, err := s.user.FindUser(ctx, username)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := VerifyPassword(password, existingUser.PasswordHash); err != nil {
		return nil, ErrInvalidCredentials
	}

	lastLoginAt := time.Now().UTC()
	if err := s.user.UpdateLastLogin(ctx, existingUser.ID, lastLoginAt); err != nil {
		return nil, err
	}
	existingUser.LastLoginAt = &lastLoginAt

	return existingUser, nil
}
