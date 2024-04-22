package types

type Type int

const (
	TypeMessage  Type = iota
	TypeJoinRoom Type = iota
)

type Body struct {
	Type     Type   `json:"type"`
	SendUser *User  `json:"send_user"`
	RoomID   uint32 `json:"room_id"`
	Data     any    `json:"data"`
}

type User struct {
	ID   uint32
	Name string
}
