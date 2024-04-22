package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(ctx *gin.Context, data ...any) {
	var item any
	if len(data) > 0 {
		item = data[0]
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": item,
	})
}

func ParamError(ctx *gin.Context, messages ...string) {
	var message string
	if len(messages) > 0 {
		message = messages[0]
	} else {
		message = "param error"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    4002,
		"message": message,
	})
}

func InternalServerError(ctx *gin.Context, messages ...string) {
	var message string
	if len(messages) > 0 {
		message = messages[0]
	} else {
		message = "internal server error"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    5001,
		"message": message,
	})
}

func UnauthorizedError(ctx *gin.Context, messages ...string) {
	var message string
	if len(messages) > 0 {
		message = messages[0]
	} else {
		message = "unauthorized error"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    4001,
		"message": message,
	})
}
