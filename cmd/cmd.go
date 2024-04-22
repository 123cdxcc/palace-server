package main

import (
	"palace-server/configs"
	"palace-server/internal/handler"
	"palace-server/internal/logic"
	"palace-server/internal/router"
)

func main() {
	conf := &configs.Config{
		AuthWhiteList: []string{
			"/user/register",
		},
	}
	userLogic := logic.NewUserLogic()
	core := logic.NewCore()
	userHandler := handler.NewUserHandler(userLogic)
	roomLogic := logic.NewRoomLogic()
	roomHandler := handler.NewRoomHandler(roomLogic)
	connectHandler := handler.NewConnectHandler(core, userLogic, roomLogic)
	go core.Run()
	router.Init(conf, userHandler, roomHandler, connectHandler)
	if err := router.Start("localhost:8081"); err != nil {
		panic(err)
	}
}
