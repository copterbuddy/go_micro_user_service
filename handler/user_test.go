package handler_test

import (
	"encoding/json"
	"main/handler"
	"main/repository"
	"main/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_User_GetallUser_Success(t *testing.T) {

	testcase_success := []repository.User{
		{
			// Id:       1,
			Email:    "cop1@test.com",
			Password: "1234",
			Name:     "Cop1",
		},
		{
			// Id:       2,
			Email:    "cop2@test.com",
			Password: "1234",
			Name:     "Cop2",
		},
	}

	expedted_mock := []service.UserResponse{
		{
			Email: "cop1@test.com",
			Name:  "Cop1",
		},
		{
			Email: "cop2@test.com",
			Name:  "Cop2",
		},
	}

	// Arrange
	t.Run("SuccessCase", func(t *testing.T) {
		expected := expedted_mock

		userService := service.NewUserServiceMock()
		userService.On("GetAll").Return(testcase_success, nil)

		userHandler := handler.NewUserHandler(userService)

		app := gin.Default()
		app.GET("/UserApi/GetAllUser", userHandler.GetAllUser)

		// Act
		req := httptest.NewRequest("GET", "/UserApi/GetAllUser", nil)
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		users := []service.UserResponse{}
		json.Unmarshal(w.Body.Bytes(), &users)
		// logs.Info("coptest : " + users)

		// Assert
		if assert.Equal(t, http.StatusOK, w.Result().StatusCode) {
			assert.Equal(t, expected, users)
		}
	})

}
