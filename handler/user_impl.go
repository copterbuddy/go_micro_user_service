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

		c.SetCookie("butterfly_cookie", res.Token, 10, "/", "localhost:8080", true, true)
		c.JSON(http.StatusOK, res.Name)
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
	issuer, ok := c.Get("Issuer")

	if !ok {
		c.JSON(401,
			gin.H{"status": "unable to bind data",
				"errorMessage": "unauthorize",
			})
		return
	}

	// issuerStr = string(issuer)
	res, err := h.userService.GetUserProfile(fmt.Sprintf("%v", issuer))
	if err != nil {
		c.JSON(401,
			gin.H{"status": "user not found",
				"errorMessage": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, res)
}
