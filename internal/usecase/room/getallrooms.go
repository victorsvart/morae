package room

import (
	"context"
	"errors"
	"morae/internal/domain/roomdomain"
	"morae/internal/dto/roomdto"
	"morae/internal/mapper/roommapper"
	"morae/internal/store/mongodb"
)

type GetAllRoomsUsecase interface {
  Execute(context.Context, *roomdto.GetRoomPaged) ([]*roomdomain.Room, error)
}

type GetAllRooms struct {
  repo mongodb.RoomRepository 
}

func (g *GetAllRooms) Execute(ctx context.Context, input *roomdto.GetRoomPaged) ([]*roomdomain.Room, error) {
  docs, err := g.repo.GetAllRooms(ctx, input.Page, input.PerPage)
  if docs == nil {
    return nil, errors.New("No documents found")
  }

  if err != nil {
    return nil, err
  }
  
  return roommapper.ToDomainSlice(docs), nil
}
