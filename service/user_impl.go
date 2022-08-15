package service

import (
	"main/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Create(email string, password string, name string) (UserResponse, error) {
	user, err := s.userRepo.Create(email, password, name)
	if err != nil {
		return UserResponse{}, err
	}

	response := UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}
	return response, nil
}

func (s userService) GetAll() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := []UserResponse{}
	for _, user := range users {
		item := UserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
		response = append(response, item)
	}
	return response, nil
}
