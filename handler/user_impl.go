package handler

import (
	"fmt"
	"main/logs"
	"main/model"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandler{userService: userService}
}

//http://localhost:8000/UserService/Ping
func (h userHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//http://localhost:8000/UserService/GetAllUser
func (h userHandler) GetAllUser(c *gin.Context) {

	users, err := h.userService.GetAll()
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusExpectationFailed, "expected error")
		return
	}

	c.JSON(http.StatusOK, users)
	return
}

//http://localhost:8000/UserService/CreateUser
func (h userHandler) CreateUser(c *gin.Context) {

	req := model.CreateUserRequest{}

	if c.ShouldBind(&req) == nil {

		res, err := h.userService.Create(req.Email, req.Password, req.Name)
		if err != nil {
			logs.Error(err)
			c.JSON(401, gin.H{"status": "cannot create"})
			return
		}

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
		return
	}

}

func (h userHandler) Login(c *gin.Context) {

	req := model.LoginRequest{}
	var err error

	if c.ShouldBind(&req) == nil {
		res, err := h.userService.Login(req)
		if err != nil {
			c.JSON(401,
				gin.H{"status": "failed",
					"errorMessage": err.Error(),
				})
			return
		}

		c.JSON(http.StatusOK, res)
		return
	} else {
		c.JSON(401,
			gin.H{"status": "unable to bind data",
				"errorMessage": err.Error(),
			})
		return
	}

}

//http://localhost:8000/UserService/GetUserProfile
func (h userHandler) GetUserProfile(c *gin.Context) {
	fmt.Println(c.Request.Header["Authorization"])
	issuer, ok := c.Get("Issuer")
	if ok == false {
		c.JSON(401,
			gin.H{"status": "unable to bind data",
				"errorMessage": "unauthorize",
			})
		return
	}

	fmt.Println(issuer)

	c.String(http.StatusOK, "pong")
}
