package logic

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"palace-server/pkg/log"
	"palace-server/types"
)

type UserConnect struct {
	userID uint32
	roomID uint32
	sc     chan types.Body
	conn   *websocket.Conn
	core   *Core
}

func NewUserConnect(userID, roomID uint32, conn *websocket.Conn, core *Core) *UserConnect {
	uc := &UserConnect{
		userID: userID,
		roomID: roomID,
		sc:     make(chan types.Body),
		conn:   conn,
		core:   core,
	}
	return uc
}

func (u *UserConnect) doReader() {
	for {
		_, data, err := u.conn.ReadMessage()
		if err != nil {
			u.core.remove <- u
			break
		}
		body := new(types.Body)
		err = json.Unmarshal(data, body)
		if err != nil {
			log.Errorf("unmarshal err: %v", err)
			continue
		}
		switch body.Type {
		case types.TypeJoinRoom:
			u.roomID = body.RoomID
		case types.TypeMessage:
			user, err := getUserByID(u.userID)
			if err != nil {
				log.Errorf("get user with id err: %v", err)
				continue
			}
			body.SendUser = &types.User{
				ID:   user.ID,
				Name: user.Name,
			}
			body.RoomID = u.roomID
			u.core.message <- *body
		default:
			log.Infof("unknown type: %v", body.Type)
		}
	}
}

func (u *UserConnect) doWrite() {
	for body := range u.sc {
		data, err := json.Marshal(&body)
		if err != nil {
			continue
		}
		_ = u.conn.WriteMessage(websocket.TextMessage, data)
	}
	u.conn.Close()
}
