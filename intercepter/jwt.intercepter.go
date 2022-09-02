package intercepter

import (
	"errors"
	"fmt"
	"main/logs"
	"main/model"

	"github.com/gin-gonic/gin"
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

func JwtVerify(token string, c *gin.Context) bool {
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
		fmt.Println()
		fmt.Printf("login success %v", myClaims.StandardClaims.Issuer)
		fmt.Println()

		c.Set("Issuer", myClaims.StandardClaims.Issuer)
	} else {
		return false
	}

	return true
}

// func JwtData(token string) interface{} {
// 	usertoken, err := jwt.ParseWithClaims(token, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		secret := viper.GetString("jwt.token_secret")
// 		return []byte(secret), nil
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	if myClaims, ok := usertoken.Claims.(*model.MyCustomClaims); ok && usertoken.Valid {
// 		fmt.Printf("login success %v", myClaims.StandardClaims.ExpiresAt)
// 		fmt.Printf("login success %v", myClaims.StandardClaims.Issuer)
// 		fmt.Println("")
// 	} else {
// 		return false
// 	}

// 	return true
// }
