package room
import "morae/internal/store/mongodb"

type RoomUsecases struct {
  GetById GetRoomByIdUsecase
  CreateRoom CreateRoomInterface 
}

func NewRoomUsecases(repo mongodb.RoomRepository) *RoomUsecases {
  return &RoomUsecases{
    GetById: &GetRoomById{repo},
    CreateRoom: &CreateRoom{repo},
  }
}
