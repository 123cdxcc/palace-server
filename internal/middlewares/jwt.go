package middlewares

import (
	"github.com/gin-gonic/gin"
	"palace-server/pkg/constant"
	"palace-server/pkg/resp"
	"palace-server/pkg/utils"
)

func VerifyAuth(authWhiteList []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestPath := ctx.Request.URL.Path
		for _, path := range authWhiteList {
			if requestPath == path {
				ctx.Next()
				return
			}
		}
		token := ctx.Request.URL.Query().Get("Authorization")
		hToken := ctx.GetHeader("Authorization")
		if hToken != "" {
			token = hToken
		}
		if token == "" {
			ctx.Abort()
			resp.UnauthorizedError(ctx)
			return
		}
		claims, err := utils.VerifyToken(token)
		if err != nil || claims == nil {
			ctx.Abort()
			resp.UnauthorizedError(ctx)
			return
		}
		ctx.Set(constant.UserIdKey, claims.ID)
		ctx.Next()
	}
}
