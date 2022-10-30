package handler

import (
	"fmt"
	"main/intercepter"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRouter(r *gin.Engine, services ...interface{}) {

	var userHandler UserHandler
	var intercepterService intercepter.Intercepter

	if services == nil {
		panic("services nil")
	}

	for _, service := range services {
		switch item := service.(type) {
		case UserHandler:
			userHandler = item
		case intercepter.Intercepter:
			intercepterService = item
		}
	}

	appVersion := viper.GetString("app.version")
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, fmt.Sprintf("app running in version : %v", appVersion))
	})

	router1 := r.Group("/UserService")
	{
		router1.POST("/CreateUser", userHandler.CreateUser)
		router1.POST("/Login", userHandler.Login)
		router1.GET("/Ping", userHandler.Ping)
		router1.GET("/TestCallService", userHandler.Ping)
	}

	router2 := r.Group("/UserService")
	router2.Use(intercepterService.GeneralInterceptor)
	{
		router2.POST("/GetUserProfile", userHandler.GetUserProfile)
		router2.GET("/GetAllUser", userHandler.GetAllUser)
	}
}
