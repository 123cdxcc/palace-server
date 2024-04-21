package model

type Room struct {
	ID      uint32
	Master  *User
	Players []*User
}
