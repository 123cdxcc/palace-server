package router

import "palace-server/internal/handler"

func initializeUserHandler() {
	userHandler := handler.NewUserHandler()
	r := router.Group("user")
	r.POST("register", userHandler.Register)
	r.POST("logout", userHandler.Logout)
}
