package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
	"palace-server/internal/logic"
	"palace-server/internal/model"
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
	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": token,
	})
	return
}

func (h *UserHandler) Logout(ctx *gin.Context) {

}
