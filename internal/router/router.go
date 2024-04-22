package router

import (
	"github.com/gin-gonic/gin"
	"palace-server/configs"
	"palace-server/internal/handler"
	"palace-server/internal/middlewares"
)

var router *gin.Engine

func Init(conf *configs.Config, userHandler *handler.UserHandler, connectHandler *handler.ConnectHandler) {
	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.VerifyAuth(conf.AuthWhiteList))
	initializeUserRouter(userHandler)
	initializeConnectRouter(connectHandler)
}

func initializeUserRouter(userHandler *handler.UserHandler) {
	r := router.Group("user")
	r.POST("register", userHandler.Register)
	r.POST("logout", userHandler.Logout)
}

func initializeConnectRouter(connectHandler *handler.ConnectHandler) {
	router.POST("connect", connectHandler.Connect)
}
