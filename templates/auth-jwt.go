package templates

var AuthJWT = `package auth

import (
	jwtgo "github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwtgo.RegisteredClaims
}`
