package logic

import (
	"context"
	"github.com/google/uuid"
	"palace-server/internal/model"
	"palace-server/pkg/errors"
)

type UserLogic struct {
}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) Register(ctx context.Context, req *model.User) (uint32, error) {
	req.ID = uuid.New().ID()
	_users = append(_users, req)
	return req.ID, nil
}

func (l *UserLogic) Logout(ctx context.Context, id uint32) error {
	index := -1
	for i, user := range _users {
		if user.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("user is not exist")
	}
	_users = append(_users[:index], _users[index+1:]...)
	return nil
}
