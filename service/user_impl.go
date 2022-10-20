package service

import (
	"errors"
	"main/intercepter"
	"main/logs"
	"main/model"
	"main/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Create(email string, password string, name string) (*model.CreateUserResponse, error) {

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.New("StatusUnprocessableEntity")
	}

	user, err := s.userRepo.Create(email, string(passwordEncrypted), name)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := &model.CreateUserResponse{
		Email: user.Email,
		Name:  user.Name,
	}
	return response, nil
}

func (s userService) GetAll() ([]model.CreateUserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := []model.CreateUserResponse{}
	for _, user := range users {
		item := model.CreateUserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
		response = append(response, item)
	}
	return response, nil
}

func (s userService) Login(req model.LoginRequest) (res *model.LoginResponse, err error) {

	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		logs.Error(err)
		return res, err
	}

	err = intercepter.CompareHashAndPassword(user.Password, req.Password)
	if err != nil {
		return res, errors.New("username or password incorrect")
	}

	// Create the Claims
	claims := model.MyCustomClaims{
		// "bar",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    strconv.FormatUint(uint64(user.ID), 10),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := viper.GetString("jwt.token_secret")
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, errors.New("unautuorize")
	}
	//end token

	res = &model.LoginResponse{
		Token: token,
		Name:  user.Name,
	}

	return res, nil
}

func (s userService) GetUserProfile(userId string) (user *model.GetUserProfileResponse, err error) {

	return nil, nil
}
