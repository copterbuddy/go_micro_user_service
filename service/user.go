package service

import "main/model"

type UserService interface {
	Create(email string, password string, name string) (user *model.CreateUserResponse, err error)
	GetAll() (users []model.CreateUserResponse, err error)
	Login(model.LoginRequest) (res *model.LoginResponse, err error)
	GetUserProfile(userID int) (user *model.GetUserProfileResponse, err error)
}
