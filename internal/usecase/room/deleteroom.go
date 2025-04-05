package room

import (
	"context"
	"morae/internal/store/mongodb"
)

type DeleteRoomUsecase interface {
  Execute(context.Context, string) error
}

type DeleteRoom struct {
  repo mongodb.RoomRepository
}

func (d *DeleteRoom) Execute(ctx context.Context, id string) error {
  if err := d.repo.DeleteRoom(ctx, id); err != nil {
    return err
  }

  return nil
}
