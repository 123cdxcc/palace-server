package handler

import (
	"github.com/gin-gonic/gin"
	"palace-server/internal/logic"
	"palace-server/pkg/log"
	"palace-server/pkg/resp"
	"palace-server/pkg/utils"
	"strconv"
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
	log.Infof("created room %v", room.ID)
}

func (h *RoomHandler) ListRoom(ctx *gin.Context) {
	rooms, err := h.roomLogic.ListRoom(ctx)
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	resp.Ok(ctx, rooms)
}

func (h *RoomHandler) ListRoomUser(ctx *gin.Context) {
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
	users, err := h.roomLogic.ListRoomUser(ctx, uint32(roomID))
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	resp.Ok(ctx, users)
	return
}
