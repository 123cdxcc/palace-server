package router

import (
	"github.com/gin-gonic/gin"
	"palace-server/configs"
	"palace-server/internal/handler"
	"palace-server/internal/middlewares"
)

var router *gin.Engine

func Start(addr ...string) error {
	return router.Run(addr...)
}

func Init(conf *configs.Config, userHandler *handler.UserHandler, roomHandler *handler.RoomHandler, connectHandler *handler.ConnectHandler) {
	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.VerifyAuth(conf.AuthWhiteList))
	initializeUserRouter(userHandler)
	initializeRoomRouter(roomHandler)
	initializeConnectRouter(connectHandler)
}

func initializeUserRouter(userHandler *handler.UserHandler) {
	r := router.Group("user")
	r.POST("register", userHandler.Register)
	r.POST("logout", userHandler.Logout)
}

func initializeRoomRouter(roomHandler *handler.RoomHandler) {
	r := router.Group("room")
	r.POST("create", roomHandler.CreateRoom)
	r.GET("list", roomHandler.ListRoom)
	r.GET("list-user", roomHandler.ListRoomUser)
}

func initializeConnectRouter(connectHandler *handler.ConnectHandler) {
	router.GET("connect", connectHandler.Connect)
}
