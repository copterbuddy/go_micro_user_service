package service

import "github.com/stretchr/testify/mock"

type userServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (m *userServiceMock) GetAll() ([]UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]UserResponse), args.Error(1)
}

func (m *userServiceMock) Create(email string, password string, name string) (UserResponse, error) {
	args := m.Called()
	return args.Get(0).(UserResponse), args.Error(1)
}
