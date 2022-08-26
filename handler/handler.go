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

	router := r.Group("/UserApi")
	{
		router.GET("/GetAllUser", userHandler.GetAllUser)
		router.POST("/CreateUser", userHandler.CreateUser)
		router.POST("/Login", userHandler.Login)
		router.GET("/ping", userHandler.Ping)
	}

	router1 := r.Group("/UserService")
	router1.Use(intercepter.GeneralInterceptor)
	{
		router1.POST("/GetUserProfile", userHandler.GetUserProfile)
	}
}
