package handler_test

import (
	"encoding/json"
	"main/handler"
	"main/model"
	"main/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ping_Success(t *testing.T) {

	//Arrange
	gin.SetMode(gin.TestMode)

	userService := service.NewUserServiceMock()
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	router := r.Group("/UserApi").GET("/ping", userHandler.Ping)
	{
		router.GET("/ping", userHandler.Ping)
	}
	res := httptest.NewRecorder()

	//Act
	req, _ := http.NewRequest("GET", "/UserApi/ping", nil)
	r.ServeHTTP(res, req)

	//Asset
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}

func Test_User_GetallUser_Success(t *testing.T) {

	want := []model.UserResponse{
		{
			Email: "cop1@test.com",
			Name:  "Cop1",
		},
		{
			Email: "cop2@test.com",
			Name:  "Cop2",
		},
	}

	gin.SetMode(gin.TestMode)

	// Arrange
	t.Run("SuccessCase", func(t *testing.T) {

		userService := service.NewUserServiceMock()
		userService.On("GetAll").Return(want, nil)

		userHandler := handler.NewUserHandler(userService)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		router := r.Group("/UserApi")
		{
			router.GET("/GetAllUser", userHandler.GetAllUser)
		}
		res := httptest.NewRecorder()

		// Act
		req, _ := http.NewRequest("GET", "/UserApi/GetAllUser", nil)
		r.ServeHTTP(res, req)

		// Assert
		users := []model.UserResponse{}
		json.Unmarshal(res.Body.Bytes(), &users)

		if !assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			t.Errorf("handler returned wrong status code: got %v want %v",
				res.Result().StatusCode, http.StatusOK)
		}

		assert.Equal(t, want, users)
	})

}

func Test_User_GetUserProfile_Success(t *testing.T) {

	give := "cop1@test.com"
	_ = give

	want := model.UserResponse{
		Email: "cop1@test.com",
		Name:  "Cop1",
	}

	gin.SetMode(gin.TestMode)

	// Arrange
	t.Run("SuccessCase", func(t *testing.T) {

		userService := service.NewUserServiceMock()
		userService.On("GetUserProfile").Return(want, nil)

		userHandler := handler.NewUserHandler(userService)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/GetUserProfile", userHandler.GetAllUser)
		w := httptest.NewRecorder()

		// Act
		req, _ := http.NewRequest("GET", "/UserApi/GetUserProfile", nil)
		r.ServeHTTP(w, req)

		// Assert
		users := []model.UserResponse{}
		json.Unmarshal(w.Body.Bytes(), &users)

		if !assert.Equal(t, http.StatusOK, w.Result().StatusCode) {
			t.Errorf("handler returned wrong status code: got %v want %v",
				w.Result().StatusCode, http.StatusOK)
		}

		assert.Equal(t, want, users)
	})

}
