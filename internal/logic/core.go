package logic

import (
	"context"
	"fmt"
	"palace-server/pkg/log"
	"palace-server/types"
)

type Core struct {
	userConnects map[*UserConnect]struct{}
	message      chan types.Body
	remove       chan *UserConnect
	add          chan *UserConnect
}

func NewCore() *Core {
	return &Core{
		userConnects: make(map[*UserConnect]struct{}),
		message:      make(chan types.Body),
		remove:       make(chan *UserConnect),
		add:          make(chan *UserConnect),
	}
}

func (c *Core) Dispose(ctx context.Context, uc *UserConnect) {
	c.add <- uc
	go uc.doWrite()
	uc.doReader()
}

func (c *Core) Quit(ctx context.Context, uc *UserConnect) {
	c.remove <- uc
}

func (c *Core) Run() {
	for {
		select {
		case body := <-c.message: // 发送消息
			for connect := range c.userConnects {
				if body.SendUser == nil {
					body.SendUser = &types.User{
						ID:   0,
						Name: "系统",
					}
				}
				if body.SendUser.ID == connect.userID { // 如果发送者和接收的用户相同就跳过
					//continue
				}
				if connect.roomID != body.RoomID { // 不属于同一个房间的跳过
					continue
				}
				connect.sc <- body
			}
		case uc := <-c.add: // 新连接
			user, err := getUserByID(uc.userID)
			if err != nil {
				log.Errorf("%v", err)
				continue
			}
			c.userConnects[uc] = struct{}{}
			body := types.Body{
				Type: types.TypeMessage,
				SendUser: &types.User{
					ID:   0,
					Name: "系统",
				},
				RoomID: uc.roomID,
				Data:   fmt.Sprintf("hello %s", user.Name),
			}
			uc.sc <- body
		case uc := <-c.remove: // 退出连接
			if _, ok := c.userConnects[uc]; ok {
				delete(c.userConnects, uc)
				close(uc.sc)
			}
		}
	}
}
