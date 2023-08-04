package auth

import (
	jwtgo "github.com/golang-jwt/jwt/v5"
)

var USER_JWT_TOKEN _JWTToken

type JWTClaims struct {
	jwtgo.RegisteredClaims
}

type _JWTToken struct {
	TokenExpTime int64
	SecretKey    string
}
