package adapters

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"

	"github.com/artyyomn/go-assignment/internal/auth"
)

type SQLiteSessionRepository struct {
	db *sql.DB
}

func NewSQLiteSessionRepository(db *sql.DB) *SQLiteSessionRepository {
	return &SQLiteSessionRepository{db: db}
}

func (r *SQLiteSessionRepository) Create(ctx context.Context, session *auth.Session) error {
	if session.ID == "" {
		id, err := newSessionID()
		if err != nil {
			return err
		}
		session.ID = id
	}

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO sessions (id, user_id, created_at, expires_at) VALUES (?, ?, ?, ?)`,
		session.ID,
		session.UserID,
		session.CreatedAt,
		session.ExpiresAt,
	)
	return err
}

func (r *SQLiteSessionRepository) Find(ctx context.Context, id string) (*auth.Session, error) {
	var session auth.Session
	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, created_at, expires_at FROM sessions WHERE id = ?`,
		id,
	).Scan(&session.ID, &session.UserID, &session.CreatedAt, &session.ExpiresAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, auth.ErrSessionNotFound
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SQLiteSessionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM sessions WHERE id = ?`, id)
	return err
}

func newSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
