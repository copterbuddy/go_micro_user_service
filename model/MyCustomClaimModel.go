package model

import "github.com/golang-jwt/jwt"

type MyCustomClaims struct {
	// Foo string `json:"foo"`
	jwt.StandardClaims
}
