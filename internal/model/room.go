package model

type Room struct {
	ID      uint32  `json:"id"`
	Master  *User   `json:"master"`
	Players []*User `json:"players"`
}
