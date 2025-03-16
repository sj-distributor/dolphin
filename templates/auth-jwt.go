package templates

var AuthJWT = `package auth

import (
	"errors"
	"time"

	jwtgo "github.com/golang-jwt/jwt/v5"
	"{{.Config.Package}}/utils"
)

// 用户token
var USER_JWT_TOKEN = JWTToken{
	TokenExpTime: config.USER_TOKEN_EXP_TIME,
	SecretKey:    config.USER_TOKEN_SECRET_KEY,
}

// 管理员token
var ADMIN_JWT_TOKEN = JWTToken{
	TokenExpTime: config.ADMIN_TOKEN_EXP_TIME,
	SecretKey:    config.ADMIN_TOKEN_SECRET_KEY,
}


type JWTClaims struct {
	jwtgo.RegisteredClaims
}

type JWTToken struct {
	TokenExpTime int64
	SecretKey    string
}

// 設置JWT
func (j *JWTToken) SetToken(str interface{}) (string, error) {
	timeNow := time.Now().Unix()
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"content": str,
		"nbf":     int64(timeNow),
		"exp":     int64(timeNow + 60*60*24*j.TokenExpTime),
	})

	return token.SignedString([]byte(j.SecretKey))
}

// 驗證JWT有效性
func (j *JWTToken) Verify(token string) error {
	_, err := j.GetTokenContent(token)
	if err != nil {
		return err
	}

	_, err = j.DecryptToken(token)
	if err != nil {
		return err
	}

	return nil
}

/**
 * token解密
 */
func (j *JWTToken) DecryptToken(token string) (map[string]interface{}, error) {
	claims, err := j.GetTokenContent(token)

	if claims == nil || err != nil {
		return map[string]interface{}{}, errors.New("Invalid Authorization")
	}

	reqData := claims.(map[string]interface{})

	return reqData, err
}

/**
 * 获取token内容
 */
func (j *JWTToken) GetTokenContent(token string) (interface{}, error) {
	if len(token) < 7 {
		return nil, errors.New("Invalid Authorization")
	}

	tokenStr := strings.Replace(token, "Bearer ", "", 1)
	claims, err := j.ParseToken(tokenStr, []byte(j.SecretKey))

	return claims["content"], err
}

/**
 * 校验token是否有效
 */
func (j *JWTToken) ParseToken(data string, key []byte) (jwtgo.MapClaims, error) {
	token, err := jwtgo.Parse(data, func(token *jwtgo.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwtgo.MapClaims)

	return claims, nil
}`
