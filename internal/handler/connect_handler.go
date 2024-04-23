package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"palace-server/internal/logic"
	"palace-server/pkg/log"
	"palace-server/pkg/resp"
	"palace-server/pkg/utils"
	"strconv"
)

type ConnectHandler struct {
	upgrader  websocket.Upgrader
	core      *logic.Core
	userLogic *logic.UserLogic
	roomLogic *logic.RoomLogic
}

func NewConnectHandler(core *logic.Core, userLogic *logic.UserLogic, roomLogic *logic.RoomLogic) *ConnectHandler {
	return &ConnectHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		core:      core,
		userLogic: userLogic,
		roomLogic: roomLogic,
	}
}

func (h *ConnectHandler) Connect(ctx *gin.Context) {
	roomIDStr, ok := ctx.GetQuery("room-id")
	if !ok {
		resp.ParamError(ctx)
		return
	}
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	_, err = h.roomLogic.GetRoom(ctx, uint32(roomID))
	if err != nil {
		resp.ParamError(ctx, err.Error())
		return
	}
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		log.Errorf("get user id err: %v", err)
		return
	}
	err = h.roomLogic.JoinRoom(ctx, userID, uint32(roomID))
	if err != nil {
		log.Errorf("join err: %v", err)
		resp.InternalServerError(ctx, err.Error())
		return
	}
	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Errorf("resp upgrade to websocket err: %v", err)
		return
	}
	uc := logic.NewUserConnect(userID, uint32(roomID), conn, h.core)
	log.Infof("user %v connect", userID)
	h.core.Dispose(ctx, uc)
	defer func() {
		h.core.Quit(ctx, uc)
		log.Infof("user %v disconnect", userID)
	}()
}
