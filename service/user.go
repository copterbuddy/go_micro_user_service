package service

import "main/model"

type UserService interface {
	Create(email string, password string, name string) (user *model.UserResponse, err error)
	GetAll() (users []model.UserResponse, err error)
	Login(model.LoginRequest) (res *model.LoginResponse, err error)
}
