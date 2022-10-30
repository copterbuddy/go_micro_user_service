package intercepter

import (
	"github.com/stretchr/testify/mock"
)

type intercepterMock struct {
	mock.Mock
}

func NewUserServiceMock() *intercepterMock {
	return &intercepterMock{}
}

func (m *intercepterMock) CompareHashAndPassword(user_password, req_password string) (err error) {
	args := m.Called()
	return args.Get(0).(error)
}
