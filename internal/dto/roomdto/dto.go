package roomdto

type RoomInput struct {
	OwnerId   uint64 `json:"ownerId"`
	Street   string `json:"street"`
	Number   uint16 `json:"number"`
	District string `json:"district"`
	State    string `json:"state"`
}
