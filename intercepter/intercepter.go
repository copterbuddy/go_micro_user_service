package intercepter

import "github.com/gin-gonic/gin"

type Intercepter interface {
	GeneralInterceptor(c *gin.Context)
	CompareHashAndPassword(user_password, req_password string) (err error)
}
