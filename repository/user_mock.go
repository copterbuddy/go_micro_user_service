package repository

import "github.com/stretchr/testify/mock"

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}

	// return userRepositoryMock{users: users}
}

func (m *userRepositoryMock) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func (m *userRepositoryMock) Create(email string, password string, name string) (*User, error) {
	args := m.Called()
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) Login() (*User, error) {
	args := m.Called()
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) GetUserByID(userID int) (*User, error) {
	args := m.Called()
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) GetUserByEmail(email string) (*User, error) {
	args := m.Called()
	return args.Get(0).(*User), args.Error(1)
}
