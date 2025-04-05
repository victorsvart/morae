package roomdomain

type Room struct {
	ID          string
	OwnerId     uint64
	Street      string
	Number      uint16
	District    string
	FullAddress string
	State       string
}
