package service_test

import (
	"main/model"
	"main/repository"
	"main/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User_GetAll_Success(t *testing.T) {

	//Arrange
	testcase_success := []repository.User{
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
	}

	expected_mock := []model.CreateUserResponse{
		{
			Email: "cop1@test.com",
			Name:  "Cop1",
		},
		{
			Email: "cop2@test.com",
			Name:  "Cop2",
		},
	}

	userRepo := repository.NewUserRepositoryMock() // Use Repo
	userRepo.On("GetAll").Return(testcase_success, nil)

	userService := service.NewUserService(userRepo) // Use Service

	//Act
	result, _ := userService.GetAll()
	expected := expected_mock

	//Assert
	assert.Equal(t, result, expected)
}
