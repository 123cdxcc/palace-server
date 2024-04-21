package handler

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(ctx *gin.Context) {
}

func (h *UserHandler) Logout(ctx *gin.Context) {

}
