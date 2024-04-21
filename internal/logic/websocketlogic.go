package logic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"palace-server/pkg/errors"
	"sync/atomic"
)

type wsConn = websocket.Conn
type wsConnHandlerFunc func(ctx context.Context, conn *wsConn)

// userID->roomID->wsConn
type websocketPool map[uint32]map[uint32]*wsConn

type WebSocketLogic struct {
	fnCount atomic.Uint64
	pool    websocketPool
}

func NewWebSocketLogic() *WebSocketLogic {
	return &WebSocketLogic{
		pool: make(map[uint32]map[uint32]*wsConn),
	}
}
func (l *WebSocketLogic) Put(ctx context.Context, userID, roomID uint32, conn *wsConn) error {
	user, ok := l.pool[userID]
	if !ok {
		user = make(map[uint32]*wsConn)
	}
	user[roomID] = conn
	l.pool[userID] = user
	return nil
}
func (l *WebSocketLogic) Take(ctx context.Context, userID, roomID uint32) (*wsConn, error) {
	user, ok := l.pool[userID]
	if !ok {
		return nil, errors.New("websocket conn not exist")
	}
	conn, ok := user[roomID]
	if !ok {
		return nil, errors.New("websocket conn not exist")
	}
	delete(user, roomID)
	l.pool[userID] = user
	return conn, nil
}

func (l *WebSocketLogic) Get(ctx context.Context, userID, roomID uint32) (*wsConn, error) {
	user, ok := l.pool[userID]
	if !ok {
		return nil, errors.New("websocket conn not exist")
	}
	conn, ok := user[roomID]
	if !ok {
		return nil, errors.New("websocket conn not exist")
	}
	return conn, nil
}

func (l *WebSocketLogic) Do(ctx context.Context, userID, roomID uint32, fn wsConnHandlerFunc) error {
	conn, err := l.Get(ctx, userID, roomID)
	if err != nil {
		return err
	}
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		l.fnCount.Add(1)
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("%+v", r)
			}
			l.fnCount.Add(-1)
			cancel()
		}()
		fn(ctx, conn)
	}()
	return nil
}
