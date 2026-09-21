package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/artyyomn/go-assignment/internal/user"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Create(ctx context.Context, user *user.User) error {
	// TODO implement the interface operations
	result, err := r.db.ExecContext(
		ctx,
		`
		INSERT INTO users (
			username,
			password_hash,
			failed_attempts,
			locked_until,
			created_at,
			last_login_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
		`,
		user.Username,
		user.PasswordHash,
		user.FailedAttempts,
		user.LockedUntil,
		user.CreatedAt,
		user.LastLoginAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id

	return nil
}

func (r *SQLiteRepository) FindUser(ctx context.Context, username string) (*user.User, error) {
	//TODO implement the inferface to find uers from username
	var u user.User

	err := r.db.QueryRowContext(
		ctx,
		`SELECT
		id,
		username,
		password_hash,
		failed_attempts,
		locked_until,
		created_at,
		last_login_at
		FROM users
		WHERE username = ?`,
		username,
	).Scan(
		&u.ID,
		&u.Username,
		&u.PasswordHash,
		&u.FailedAttempts,
		&u.LockedUntil,
		&u.CreatedAt,
		&u.LastLoginAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, user.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *SQLiteRepository) UpdateLastLogin(ctx context.Context, userID int64, lastLoginAt time.Time) error {
	_, err := r.db.ExecContext(
		ctx,
		`UPDATE users SET last_login_at = ? WHERE id = ?`,
		lastLoginAt,
		userID,
	)
	return err
}
