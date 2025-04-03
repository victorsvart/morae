package postgres

import (
	"database/sql"
)

type PostgresStorage struct {
	Users UserRepository
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		Users: &UserStore{db},
	}
}
