package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserEntity struct {
	ID           uint64
	FullName     string
	EmailAddress string
	Password     string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context, ue *UserEntity) error {
	query := `SELECT INTO mre_users (fullName, emaillAddress, password) 
            VALUES ($1, $2, $3)
            RETURNING id, created_at`

	err := us.db.QueryRowContext(ctx, query, ue.FullName, ue.EmailAddress, ue.Password).Scan(
		&ue.ID,
		&ue.CreatedAt,
	)
	if err != nil {
		return errors.New("Error inserting user into user")
	}

	return nil
}
