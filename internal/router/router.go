package router

import (
	"github.com/gin-gonic/gin"
	"palace-server/configs"
	"palace-server/internal/middlewares"
)

var router *gin.Engine

func Init(conf *configs.Config) {
	router = gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.VerifyAuth(conf.AuthWhiteList))
}
