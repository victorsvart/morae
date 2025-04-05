// Package postgres provides the PostgreSQL implementation of the user repository.
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
)

// UserEntity represents a user record in the database.
type UserEntity struct {
	ID           uint64
	FullName     string
	EmailAddress string
	Password     string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

// UserRepository defines the contract for user persistence.
type UserRepository interface {
	GetByID(context.Context, uint64) (*UserEntity, error)
	Create(context.Context, *UserEntity) error
	Update(context.Context, *UserEntity) error
	Delete(context.Context, uint64) error
	List(context.Context) ([]*UserEntity, error)
	FindByEmail(context.Context, string) (*UserEntity, error)
}

// UserStore implements the UserRepository using PostgreSQL.
type UserStore struct {
	db *sql.DB
}

// GetByID retrieves a user by ID.
func (us *UserStore) GetByID(ctx context.Context, id uint64) (*UserEntity, error) {
	query, args, err := qb.
		Select("id", "full_name", "email_address").
		From("mre_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user UserEntity
	err = us.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.FullName,
		&user.EmailAddress,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Create inserts a new user if the email is not already in use.
func (us *UserStore) Create(ctx context.Context, ue *UserEntity) error {
	emailInUse, err := us.emailExists(ctx, ue.EmailAddress)
	if err != nil {
		return err
	}
	if emailInUse != "" {
		return errors.New("email already in use")
	}

	query, args, err := qb.
		Insert("mre_users").
		Columns("full_name", "email_address", "password").
		Values(ue.FullName, ue.EmailAddress, ue.Password).
		Suffix("RETURNING id, full_name, email_address, created_at").
		ToSql()
	if err != nil {
		return err
	}

	return us.db.QueryRowContext(ctx, query, args...).Scan(
		&ue.ID,
		&ue.FullName,
		&ue.EmailAddress,
		&ue.CreatedAt,
	)
}

// Update modifies user details.
func (us *UserStore) Update(ctx context.Context, ue *UserEntity) error {
	valid, err := us.validID(ctx, ue.ID)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("invalid ID")
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

	query, args, err := qb.
		Update("mre_users").
		Set("full_name", squirrel.Expr("COALESCE(?, full_name)", fullName)).
		Set("email_address", squirrel.Expr("COALESCE(?, email_address)", emailAddress)).
		Set("password", squirrel.Expr("COALESCE(?, password)", password)).
		Where(squirrel.Eq{"id": ue.ID}).
		Suffix("RETURNING id, full_name, email_address, created_at, updated_at").
		ToSql()
	if err != nil {
		return err
	}

	return us.db.QueryRowContext(ctx, query, args...).Scan(
		&ue.ID,
		&ue.FullName,
		&ue.EmailAddress,
		&ue.CreatedAt,
		&ue.UpdatedAt,
	)
}

// Delete removes a user by ID.
func (us *UserStore) Delete(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("invalid ID provided")
	}

	query, args, err := qb.
		Delete("mre_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := us.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

// List retrieves all users.
// TODO: Implement user listing.
func (us *UserStore) List(_ context.Context) ([]*UserEntity, error) {
	return nil, nil
}

// FindByEmail returns a user by their email.
func (us *UserStore) FindByEmail(ctx context.Context, emailAddress string) (*UserEntity, error) {
	query, args, err := qb.
		Select("id", "full_name", "email_address", "password").
		From("mre_users").
		Where(squirrel.Eq{"email_address": emailAddress}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var userEntity UserEntity
	err = us.db.QueryRowContext(ctx, query, args...).Scan(
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

// emailExists checks if an email is already in use.
func (us *UserStore) emailExists(ctx context.Context, emailAddress string) (string, error) {
	if emailAddress == "" {
		return "", errors.New("email cannot be empty")
	}

	query, args, err := qb.
		Select("email_address").
		From("mre_users").
		Where(squirrel.Eq{"email_address": emailAddress}).
		ToSql()
	if err != nil {
		return "", err
	}

	var existing string
	err = us.db.QueryRowContext(ctx, query, args...).Scan(&existing)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		log.Println("emailExists error:", err)
		return "", err
	}

	return existing, nil
}

// validID checks if a user ID exists in the database.
func (us *UserStore) validID(ctx context.Context, id uint64) (bool, error) {
	if id == 0 {
		return false, errors.New("invalid ID")
	}

	query := `SELECT EXISTS(SELECT 1 FROM mre_users WHERE id = $1)`
	var exists bool
	err := us.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
