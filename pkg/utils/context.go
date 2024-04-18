package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"palace-server/pkg/constant"
)

func GetUserID(ctx *gin.Context) (uint32, error) {
	value, ok := ctx.Get(constant.UserIdKey)
	if !ok {
		return 0, errors.New("no user id")
	}
	id, ok := value.(uint32)
	if !ok {
		return 0, errors.New("user id value not uint32")
	}
	return id, nil
}
