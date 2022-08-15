package handler

import (
	"main/logs"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandler{userService: userService}
}

//http://localhost:8000/UserApi/GetAllUser
func (h userHandler) GetAllUser(c *gin.Context) {

	users, err := h.userService.GetAll()
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusExpectationFailed, "expected error")
	}

	c.JSON(http.StatusOK, users)
}

//http://localhost:8000/UserApi/CreateUser
func (h userHandler) CreateUser(c *gin.Context) {

	req := CreateUserReq{}
	res := service.UserResponse{}
	var err error

	if c.ShouldBind(&req) == nil {
		res, err = h.userService.Create(req.Email, req.Password, req.Name)
		if err != nil {
			logs.Error(err)
			c.JSON(401, gin.H{"status": "cannot create"})
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

	c.JSON(http.StatusOK, res)
}

type CreateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
