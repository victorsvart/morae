package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	Update(context.Context, *UserEntity) error
	Delete(context.Context, uint64) error
	List(context.Context) ([]*UserEntity, error)
	FindByEmail(ctx context.Context, emailAddress string) (*UserEntity, error)
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) GetById(ctx context.Context, id uint64) (*UserEntity, error) {
	query := "SELECT id, full_name, email_address FROM mre_users where id = $1"
	var user UserEntity
	err := us.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FullName, &user.EmailAddress)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserStore) Create(ctx context.Context, ue *UserEntity) error {
	emailInUsage, err := us.emailExists(ctx, ue.EmailAddress)

	if err != nil {
		return err
	}

	if emailInUsage != "" {
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

func (us *UserStore) Update(ctx context.Context, ue *UserEntity) error {
	idIsValid, err := us.validId(ctx, ue.ID)
	if err != nil {
		return err
	}

	if !idIsValid {
		return errors.New("Id is not valid")
	}

	var fullName, emailAddress, password *string
	if ue.FullName != "" {
		fullName = &ue.FullName
	}
	if ue.EmailAddress != "" {
		emailAddress = &ue.EmailAddress
	}
	if ue.Password != "" {
		password = &ue.Password
	}

	query := `UPDATE mre_users SET 
              full_name = COALESCE($1, full_name), 
              email_address = COALESCE($2, email_address), 
              password = COALESCE($3, password) 
            WHERE id = $4
            RETURNING id, full_name, email_address, created_at, updated_at`

	err = us.db.QueryRowContext(ctx, query, fullName, emailAddress, password, ue.ID).Scan(
		&ue.ID,
		&ue.FullName,
		&ue.EmailAddress,
		&ue.CreatedAt,
		&ue.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) Delete(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("invalid id provided")
	}

	query := `DELETE FROM mre_users WHERE id = $1`
	result, err := us.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("no user found with the given id")
	}

	return nil
}

func (us *UserStore) List(ctx context.Context) ([]*UserEntity, error) {
	return nil, nil
}

func (us *UserStore) FindByEmail(ctx context.Context, emailAddress string) (*UserEntity, error) {
	query := `SELECT id, full_name, email_address, password FROM mre_users WHERE email_address = $1`
	var userEntity UserEntity
	err := us.db.QueryRowContext(ctx, query, emailAddress).Scan(
		&userEntity.ID,
		&userEntity.FullName,
		&userEntity.EmailAddress,
		&userEntity.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &userEntity, nil
}

func (us *UserStore) emailExists(ctx context.Context, emailAddress string) (string, error) {
	if emailAddress == "" {
		return "", errors.New("Email cannot be empty")
	}

	query := `SELECT email_address FROM mre_users WHERE email_address = $1`
	var email_address string
	err := us.db.QueryRowContext(ctx, query, emailAddress).Scan(&email_address)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}

		return "", err
	}

	return email_address, nil
}

func (us *UserStore) validId(ctx context.Context, id uint64) (bool, error) {
	if id == 0 {
		return false, errors.New("Id is invalid")
	}

	query := `SELECT EXISTS(SELECT 1 FROM mre_users WHERE id = $1)`
	var exists bool
	err := us.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
