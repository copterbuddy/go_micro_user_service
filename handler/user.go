package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Ping(c *gin.Context)
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetUserProfile(c *gin.Context)
}
