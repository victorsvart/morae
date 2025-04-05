package room

import (
	"context"
	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

type UpdateRoomUsecase interface {
	Execute(context.Context, *roomdto.RoomDto) (roomdomain.Room, error)
}

type UpdateRoom struct {
	repo mongodb.RoomRepository
}

func (u *UpdateRoom) Execute(ctx context.Context, input *roomdto.RoomDto) (roomdomain.Room, error) {
	document, err := roommapper.FromDto(input)
  if err != nil  {
    return roomdomain.Room{}, err
  }

  err = u.repo.UpdateRoom(ctx, &document)
  if err != nil {
    return roomdomain.Room{}, nil
  }

  return roommapper.ToDomain(&document), nil
}
