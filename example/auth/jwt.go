package auth

import (
	jwtgo "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwtgo.RegisteredClaims
}
