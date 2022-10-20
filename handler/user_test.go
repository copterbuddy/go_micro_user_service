package handler_test

import (
	"bytes"
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

func Test_User_Ping_Success(t *testing.T) {

	t.Run("PingSuccesss", func(t *testing.T) {
		//Arrange
		gin.SetMode(gin.TestMode)

		userService := service.NewUserServiceMock()
		userHandler := handler.NewUserHandler(userService)

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)

		router := r.Group("/UserService")
		{
			router.GET("/Ping", userHandler.Ping)
		}
		// res := httptest.NewRecorder()

		//Act
		req, _ := http.NewRequest("GET", "/UserService/Ping", nil)
		r.ServeHTTP(res, req)

		//Asset
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "pong", res.Body.String())

	})
}

func Test_User_Login_Success(t *testing.T) {

	t.Run("Successs", func(t *testing.T) {
		//Arrange
		gin.SetMode(gin.TestMode)

		given := model.LoginRequest{
			Email:    "copemail@test.com",
			Password: "1234",
		}

		want := &model.LoginResponse{
			Name: "testName",
		}

		userService := service.NewUserServiceMock()
		userHandler := handler.NewUserHandler(userService)

		var loginRequestByte []byte
		loginRequestByte, err := json.Marshal(given)
		if err != nil {
			t.Error(err)
		}

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)

		router := r.Group("/UserService")
		{
			router.POST("/Login", userHandler.Login)
		}

		userService.On("Login").Return(want, nil)

		//Act
		req, _ := http.NewRequest(http.MethodPost, "/UserService/Login", bytes.NewBuffer(loginRequestByte))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		// resStruct := &model.LoginResponse{}
		name := ""
		json.NewDecoder(res.Body).Decode(&name)

		//Asset

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, want.Name, name)

	})
}

func Test_User_GetAllUser_Success(t *testing.T) {

	want := []model.CreateUserResponse{
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

	t.Run("SuccessCase", func(t *testing.T) {

		// Arrange
		userService := service.NewUserServiceMock()
		userService.On("GetAll").Return(want, nil)

		userHandler := handler.NewUserHandler(userService)

		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()

		_, r := gin.CreateTestContext(res)

		{
			r.POST("/UserService/GetAllUser", userHandler.GetAllUser)
		}

		// Act
		req, _ := http.NewRequest(http.MethodPost, "/UserService/GetAllUser", nil)
		r.ServeHTTP(res, req)

		// Assert
		users := []model.CreateUserResponse{}
		json.Unmarshal(res.Body.Bytes(), &users)

		if !assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			t.Errorf("handler returned wrong status code: got %v want %v",
				res.Result().StatusCode, http.StatusOK)
		}

		assert.Equal(t, want, users)
	})

}

func Test_User_GetUserProfile_Success(t *testing.T) {

	given := "cop1@test.com"
	_ = given

	want := &model.GetUserProfileResponse{

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

		res := httptest.NewRecorder()
		_, r := gin.CreateTestContext(res)
		r.Use(func(c *gin.Context) {
			c.Set("Issuer", "1")
		})

		r.POST("/UserService/GetUserProfile", userHandler.GetUserProfile)

		// Act
		req, _ := http.NewRequest(http.MethodPost, "/UserService/GetUserProfile", nil)
		r.ServeHTTP(res, req)

		// Assert
		users := &model.GetUserProfileResponse{}
		json.Unmarshal(res.Body.Bytes(), &users)

		if !assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			t.Errorf("handler returned wrong status code: got %v want %v",
				res.Result().StatusCode, http.StatusOK)
		}

		assert.Equal(t, want, users)
	})

}
