package postgres

import (
	"context"
	"database/sql"
)

type PostgresStorage struct {
	Users interface {
		Create(context.Context, *UserEntity) error
	}
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		Users: &UserStore{db},
	}
}
