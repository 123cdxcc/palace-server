package logic

import (
	"palace-server/internal/model"
	"palace-server/pkg/errors"
)

var _users = make([]*model.User, 0, 10)
var _rooms = make([]*model.Room, 0, 10)

func getUserByID(id uint32) (*model.User, error) {
	for _, user := range _users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user is not exist")
}

func getRoomByID(id uint32) (*model.Room, error) {
	for _, room := range _rooms {
		if room.ID == id {
			return room, nil
		}
	}
	return nil, errors.New("room is not exist")
}
