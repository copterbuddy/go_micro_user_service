package service

import (
	"main/model"

	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (m *userServiceMock) GetAll() ([]model.CreateUserResponse, error) {
	args := m.Called()
	return args.Get(0).([]model.CreateUserResponse), args.Error(1)
}

func (m *userServiceMock) Create(email string, password string, name string) (*model.CreateUserResponse, error) {
	args := m.Called()
	return args.Get(0).(*model.CreateUserResponse), args.Error(1)
}

func (m *userServiceMock) Login(model.LoginRequest) (res *model.LoginResponse, err error) {
	args := m.Called()
	return args.Get(0).(*model.LoginResponse), args.Error(1)
}

func (m *userServiceMock) GetUserProfile(userId string) (user *model.GetUserProfileResponse, err error) {
	args := m.Called()
	return args.Get(0).(*model.GetUserProfileResponse), args.Error(1)
}
