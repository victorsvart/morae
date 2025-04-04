package user

import (
	"context"
	"errors"
	"morae/internal/store/postgres"
)

type DeleteUserUsecase interface {
	Execute(context.Context, uint64) error
}

type Delete struct {
	repo postgres.UserRepository
}

func (d *Delete) Execute(ctx context.Context, id uint64) error {
	if id == 0 {
		return errors.New("Id is invalid")
	}

	err := d.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
