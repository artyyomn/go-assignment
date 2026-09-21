package user

import (
	"context"
	"database/sql"

	"github.com/artyyomn/go-assignment/internal/user"
)

type SQLiteRepository struct{
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository{
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Create(ctx context.Context, u *user.User) error {
	// TODO implement the interface operations
    return nil
}

func (r *SQLiteRepository) Find(ctx context.Context, username string) (*user.User, error){
	//TODO implement the inferface to find uers from username
	return nil,nil
}
