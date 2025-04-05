package room

import (
	"context"
	"errors"
	"morae/internal/store/mongodb"
)

type DeleteRoomUsecase interface {
  Execute(context.Context, string) error
}

type DeleteRoom struct {
  repo mongodb.RoomRepository
}

func (d *DeleteRoom) Execute(ctx context.Context, id string) error {
  if id == "" {
    return errors.New("Id cannot be empty")
  }

  if err := d.repo.DeleteRoom(ctx, id); err != nil {
    return err
  }

  return nil
}
