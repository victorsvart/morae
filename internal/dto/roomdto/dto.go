package roomdto

type RoomDto struct {
  ID          string `json:"id"`
  OwnerId     uint64 `json:"ownerId"`
  Street      string `json:"street"`
  Number      uint16 `json:"number"`
  District    string `json:"district"`
  State       string `json:"state"`
}

type RoomInput struct {
	OwnerId  uint64 `json:"ownerId"`
	Street   string `json:"street"`
	Number   uint16 `json:"number"`
	District string `json:"district"`
	State    string `json:"state"`
}

type GetRoomPaged struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"perPage"`
}
