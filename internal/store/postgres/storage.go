package postgres

import (
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
)

// postgres specific stuff. Sets all '?' query parameters as $.
var qb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

// Storage holds all PostgreSQL-backed repository implementations.
// It serves as a centralized struct to access various data sources via Postgres.
type Storage struct {
	Users UserRepository
}

// NewPostgresStorage initializes a new PostgresStorage instance,
// setting up all PostgreSQL-backed repositories using the provided database connection.
func NewPostgresStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: &UserStore{db},
	}
}

// LogQuery prints the SQL query and its arguments to the log.
// Useful for debugging database queries.
func LogQuery(sql string, args []any) {
	log.Printf("SQL: %s\nARGS: %v\n", sql, args)
}
