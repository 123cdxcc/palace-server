package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"palace-server/internal/logic"
	"palace-server/pkg/log"
	"palace-server/pkg/utils"
)

type ConnectHandler struct {
	upgrader websocket.Upgrader
	core     *logic.Core
}

func NewConnectHandler(core *logic.Core) *ConnectHandler {
	return &ConnectHandler{
		upgrader: websocket.Upgrader{},
		core:     core,
	}
}

func (h *ConnectHandler) Connect(ctx *gin.Context) {
	conn, err := h.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Errorf("http upgrade to websocket err: %v", err)
		return
	}
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		log.Errorf("get user id err: %v", err)
		return
	}
	uc := logic.NewUserConnect(userID, conn, h.core)
	h.core.Dispose(ctx, uc)
	defer h.core.Quit(ctx, uc)
}
