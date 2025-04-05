package user

import (
	"context"
	"errors"
	"morae/internal/store/postgres"
)

// DeleteUserUsecase defines the interface for deleting a user by ID.
type DeleteUserUsecase interface {
	Execute(context.Context, uint64) error
}

// Delete implements the DeleteUserUsecase using a Postgres repository.
type Delete struct {
	repo postgres.UserRepository
}

// Execute deletes a user by their ID. Returns an error if the ID is invalid or the operation fails.
func (d *Delete) Execute(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("id is invalid")
	}

	err := d.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
