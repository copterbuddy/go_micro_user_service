package intercepter

import (
	"errors"
	"fmt"
	"main/logs"
	"main/model"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(user_password, req_password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user_password), []byte(req_password))
	if err != nil {
		logs.Error(err)
		return errors.New("username or password incurrect")
	}
	return nil
}

func JwtVerify(token string) bool {
	usetoken, err := jwt.ParseWithClaims(token, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := viper.GetString("jwt.token_secret")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	if myClaims, ok := usetoken.Claims.(*model.MyCustomClaims); ok && usetoken.Valid {
		fmt.Printf("login success %v", myClaims.StandardClaims.ExpiresAt)
		fmt.Println("")
	} else {
		return false
	}

	return true
}
