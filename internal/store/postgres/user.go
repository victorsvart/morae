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

type UserRepository interface {
	GetById(context.Context, uint64) (*UserEntity, error)
	Create(context.Context, *UserEntity) error
	List(context.Context) ([]*UserEntity, error)
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) GetById(ctx context.Context, id uint64) (*UserEntity, error) {
	query := "SELECT id, fullName, emailAddress, FROM mre_users where id = $1"
	var user UserEntity
	err := us.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FullName, &user.EmailAddress)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserStore) Create(ctx context.Context, ue *UserEntity) error {
	emailInUsage, err := us.EmailExists(ctx, ue.EmailAddress)

	if err != nil {
		return err
	}

	if emailInUsage {
		return errors.New("Email already in usage")

	}

	query := `INSERT INTO mre_users (full_name, email_address, password) 
            VALUES ($1, $2, $3) RETURNING id, full_name, email_address, created_at`

	err = us.db.QueryRowContext(ctx, query, ue.FullName, ue.EmailAddress, ue.Password).Scan(
		&ue.ID,
		&ue.FullName,
		&ue.EmailAddress,
		&ue.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) List(ctx context.Context) ([]*UserEntity, error) {
	return nil, nil
}

func (us *UserStore) EmailExists(ctx context.Context, emailAddress string) (bool, error) {
	if emailAddress == "" {
		return false, errors.New("Email cannot be empty")
	}

	query := `SELECT EXISTS(SELECT 1 FROM mre_users WHERE email_address = $1)`
	var exists bool
	err := us.db.QueryRowContext(ctx, query, emailAddress).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
