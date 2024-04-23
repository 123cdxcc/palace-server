package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"palace-server/internal/logic"
	"palace-server/internal/model"
	"palace-server/pkg/log"
	"palace-server/pkg/resp"
	"palace-server/pkg/utils"
)

type UserHandler struct {
	userLogic *logic.UserLogic
}

func NewUserHandler(userLogic *logic.UserLogic) *UserHandler {
	return &UserHandler{
		userLogic: userLogic,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	reqData, err := ctx.GetRawData()
	if err != nil {
		resp.ParamError(ctx)
		return
	}
	nameJson := gjson.Get(string(reqData), "name")
	if !nameJson.Exists() {
		resp.ParamError(ctx)
		return
	}
	uid, err := h.userLogic.Register(ctx, &model.User{
		Name: nameJson.String(),
	})
	if err != nil {
		resp.ParamError(ctx)
		return
	}
	token, err := utils.GenToken(uid)
	if err != nil {
		resp.InternalServerError(ctx)
		return
	}
	resp.Ok(ctx, gin.H{
		"uid":   uid,
		"token": token,
	})
	log.Infof("%v register", uid)
	return
}

func (h *UserHandler) Logout(ctx *gin.Context) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		resp.UnauthorizedError(ctx)
		return
	}
	err = h.userLogic.Logout(ctx, userID)
	if err != nil {
		resp.UnauthorizedError(ctx)
		return
	}
	resp.Ok(ctx)
	log.Infof("%v logout", userID)
}
