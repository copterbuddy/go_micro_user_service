package service

import (
	"errors"
	"main/logs"
	"main/model"
	"main/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Create(email string, password string, name string) (*model.UserResponse, error) {

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.New("StatusUnprocessableEntity")
	}

	user, err := s.userRepo.Create(email, string(passwordEncrypted), name)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := &model.UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}
	return response, nil
}

func (s userService) GetAll() ([]model.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := []model.UserResponse{}
	for _, user := range users {
		item := model.UserResponse{
			Email: user.Email,
			Name:  user.Name,
		}
		response = append(response, item)
	}
	return response, nil
}

const jwtSecret = "mySecret"

func (s userService) Login(req model.LoginRequest) (res *model.LoginResponse, err error) {

	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		logs.Error(err)
		return res, err
	}

	// passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	// if err != nil {
	// 	return nil, errors.New("StatusUnprocessableEntity")
	// }

	// passCompare := string(passwordEncrypted)
	// if passCompare != user.Password {
	// 	return res, errors.New("username or password incurrect")
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New("username or password incurrect")
	}

	//TODO: gen token
	cliams := jwt.StandardClaims{
		Issuer:    strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, errors.New("unautuorize")
	}
	//end token

	res = &model.LoginResponse{
		Token: token,
	}
	return res, nil
}
