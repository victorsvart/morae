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
	query, args, err := qb.
		Select("id", "full_name", "email_address").
		From("mre_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var user UserEntity
	err = us.db.
		QueryRowContext(ctx, query, args...).
		Scan(
			&user.ID,
			&user.FullName,
			&user.EmailAddress,
		)
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

	query, args, err := qb.
		Insert("mre_users").
		Columns("full_name", "email_address", "password").
		Values(ue.FullName, ue.EmailAddress, ue.Password).
		Suffix("RETURNING id, full_name, email_address, created_at").
		ToSql()

	err = us.db.QueryRowContext(ctx, query, args...).Scan(
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

	query, args, err := qb.
		Update("mre_users").
		Set("full_name", squirrel.Expr("COALESCE(?, full_name)", fullName)).
		Set("email_address", squirrel.Expr("COALESCE(?, email_address)", emailAddress)).
		Set("password", squirrel.Expr("COALESCE(?, password)", password)).
		Where(squirrel.Eq{"id": ue.ID}).
		Suffix("RETURNING id, full_name, email_address, created_at, updated_at").
		ToSql()

	err = us.db.QueryRowContext(ctx, query, args...).Scan(
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

	query, args, err := qb.
		Delete("mre_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	result, err := us.db.ExecContext(ctx, query, args...)
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
	query, args, err := qb.
		Select("id", "full_name", "email_address", "password").
		From("mre_users").
		Where(squirrel.Eq{"email_address": emailAddress}).
		ToSql()

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

func (us *UserStore) emailExists(ctx context.Context, emailAddress string) (string, error) {
	if emailAddress == "" {
		return "", errors.New("Email cannot be empty")
	}

	query, args, err := qb.
		Select("email_address").
		From("mre_users").
		Where(squirrel.Eq{"email_address": emailAddress}).
		ToSql()

	var email_address string

	err = us.db.QueryRowContext(ctx, query, args...).Scan(&email_address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		log.Println("HERE")
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
