//go:build integration

package handler_test

import (
	"encoding/json"
	"main/handler"
	"main/model"
	"main/repository"
	"main/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll_Success_Integration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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

		userRepo := repository.NewUserRepositoryMock()
		userRepo.On("GetAll").Return([]repository.User{
			{
				Email:    "cop1@test.com",
				Password: "1234",
				Name:     "Cop1",
			},
			{
				Email:    "cop2@test.com",
				Password: "1234",
				Name:     "Cop2",
			},
		}, nil)

		userService := service.NewUserService(userRepo)
		userHandler := handler.NewUserHandler(userService)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		router := r.Group("/UserApi")
		{
			router.GET("/GetAllUser", userHandler.GetAllUser)
		}
		res := httptest.NewRecorder()

		//Act
		req, _ := http.NewRequest("GET", "/UserApi/GetAllUser", nil)
		r.ServeHTTP(res, req)

		//Asset
		users := []model.CreateUserResponse{}
		json.Unmarshal(res.Body.Bytes(), &users)

		if !assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			t.Errorf("handler returned wrong status code: got %v want %v",
				res.Result().StatusCode, http.StatusOK)
		}

		assert.Equal(t, want, users)
	})
}
