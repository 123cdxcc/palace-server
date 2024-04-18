package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"palace-server/pkg/constant"
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
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "No Authorized",
			})
			return
		}
		claims, err := utils.VerifyToken(token)
		if err != nil || claims == nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "No Authorized",
			})
			return
		}
		ctx.Set(constant.UserIdKey, claims.ID)
		ctx.Next()
	}
}
