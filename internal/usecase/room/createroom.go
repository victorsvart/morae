package room

import (
	"context"
	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

type CreateRoomInterface interface {
	Execute(context.Context, *roomdto.RoomInput) (roomdomain.Room, error)
}

type CreateRoom struct {
	repo mongodb.RoomRepository
}

func (c *CreateRoom) Execute(ctx context.Context, input *roomdto.RoomInput) (roomdomain.Room, error) {
	document := roommapper.FromInput(input)
	err := c.repo.CreateRoom(ctx, &document)
	if err != nil {
		return roomdomain.Room{}, nil
	}

	domain := roommapper.ToDomain(&document)
	return domain, nil
}
