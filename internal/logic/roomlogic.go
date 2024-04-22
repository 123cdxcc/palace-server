package logic

import (
	"context"
	"github.com/google/uuid"
	"palace-server/internal/model"
	"palace-server/pkg/errors"
)

type RoomLogic struct {
}

func NewRoomLogic() *RoomLogic {
	return &RoomLogic{}
}

func (l *RoomLogic) CreateRoom(ctx context.Context, userID uint32) (*model.Room, error) {
	user, err := getUserByID(userID)
	if err != nil {
		return nil, err
	}
	room := &model.Room{
		ID:      uuid.New().ID(),
		Master:  user,
		Players: append(make([]*model.User, 0, 5), user),
	}
	_rooms = append(_rooms, room)
	return room, nil
}

func (l *RoomLogic) JoinRoom(ctx context.Context, userID, roomID uint32) error {
	user, err := getUserByID(userID)
	if err != nil {
		return err
	}
	room, err := getRoomByID(roomID)
	if err != nil {
		return err
	}
	for _, player := range room.Players {
		if player.ID == userID {
			return nil
		}
	}
	room.Players = append(room.Players, user)
	_rooms = append(_rooms, room)
	return nil
}

func (l *RoomLogic) ListRoom(ctx context.Context) ([]*model.Room, error) {
	return _rooms, nil
}

func (l *RoomLogic) ExitRoom(ctx context.Context, roomID, userID uint32) error {
	room, err := getRoomByID(roomID)
	if err != nil {
		return err
	}
	index := -1
	for i, player := range room.Players {
		if player.ID == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("user not in room")
	}
	room.Players = append(room.Players[:index], room.Players[index+1:]...)
	return nil
}

func (l *RoomLogic) GetRoom(ctx context.Context, roomID uint32) (*model.Room, error) {
	return getRoomByID(roomID)
}
