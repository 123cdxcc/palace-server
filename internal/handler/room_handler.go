package handler

import (
	"github.com/gin-gonic/gin"
	"palace-server/internal/logic"
	"palace-server/pkg/resp"
	"palace-server/pkg/utils"
)

type RoomHandler struct {
	roomLogic *logic.RoomLogic
}

func NewRoomHandler(roomLogic *logic.RoomLogic) *RoomHandler {
	return &RoomHandler{
		roomLogic: roomLogic,
	}
}

func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.UnauthorizedError(ctx)
		return
	}
	room, err := h.roomLogic.CreateRoom(ctx, userID)
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	resp.Ok(ctx, room)
}

func (h *RoomHandler) ListRoom(ctx *gin.Context) {
	rooms, err := h.roomLogic.ListRoom(ctx)
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	resp.Ok(ctx, rooms)
}
