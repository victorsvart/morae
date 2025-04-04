package postgres

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
)

// postgres specific stuff. Sets all '?' parameters as $.
var qb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type PostgresStorage struct {
	Users UserRepository
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		Users: &UserStore{db},
	}
}

func LogQuery(sql string, args []any) {
	log.Printf("SQL: %s\nARGS: %v\n", sql, args)
}
