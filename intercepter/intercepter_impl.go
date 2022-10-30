package intercepter

import (
	"errors"
	"fmt"
	"main/logs"
	"main/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type intercepter struct {
}

func NewInterceptor() Intercepter {
	return intercepter{}
}

// GeneralInterceptor implements Intercepter
func (intercepter) GeneralInterceptor(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if len(token) <= len("Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	parts := strings.SplitN(token, " ", 2)
	if parts[0] != "Bearer" {
		c.JSON(401, &gin.H{
			"code":  401,
			"error": "Invalid Authorization header",
		})
		c.Abort()
		return
	}
	userId, isOk := JwtVerify(parts[1])
	if !isOk || userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Set("userId", userId)

	c.Next()
}

func (intercepter) CompareHashAndPassword(user_password, req_password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user_password), []byte(req_password))
	if err != nil {
		logs.Error(err)
		return errors.New("username or password incurrect")
	}
	return nil
}

func JwtVerify(token string) (userId int, isOk bool) {
	usertoken, err := jwt.ParseWithClaims(
		token,
		&model.MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			secret := viper.GetString("jwt.token_secret")
			return []byte(secret), nil
		})

	if err != nil {
		fmt.Println(err)
		return 0, false
	}

	myClaims, ok := usertoken.Claims.(*model.MyCustomClaims)
	if !ok || !usertoken.Valid {
		return 0, false
	}

	return myClaims.UserId, true
}
