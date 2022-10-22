package handler

import (
	"main/intercepter"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, handlers ...interface{}) {

	var userHandler UserHandler

	if handlers == nil {
		panic("handler nil")
	}

	for _, handler := range handlers {
		switch v := handler.(type) {
		case UserHandler:
			userHandler = v

		}
	}

	router1 := r.Group("/UserService")
	{
		router1.POST("/CreateUser", userHandler.CreateUser)
		router1.POST("/Login", userHandler.Login)
		router1.GET("/Ping", userHandler.Ping)
		router1.GET("/TestCallService", userHandler.Ping)
	}

	router2 := r.Group("/UserService")
	router2.Use(intercepter.GeneralInterceptor)
	{
		router2.POST("/GetUserProfile", userHandler.GetUserProfile)
		router2.GET("/GetAllUser", userHandler.GetAllUser)
	}
}
