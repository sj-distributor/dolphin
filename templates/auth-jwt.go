package templates

var AuthJWT = `package auth

import (
	"errors"
	"time"

	jwtgo "github.com/golang-jwt/jwt/v5"
	"{{.Config.Package}}/utils"
)

var USER_JWT_TOKEN _JWTToken

type JWTClaims struct {
	jwtgo.RegisteredClaims
}

type _JWTToken struct {
	TokenExpTime int64
	SecretKey    string
}

// 設置JWT
func (j *_JWTToken) SetToken(str interface{}) string {
	timeNow := time.Now().Unix()
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{
		"content": str,
		"nbf":     int64(timeNow),
		"exp":     int64(timeNow + 60*60*24*j.TokenExpTime),
	})

	ss, _ := token.SignedString([]byte(j.SecretKey))
	return ss
}

// 驗證JWT有效性
func (j *_JWTToken) Verify(token string) error {
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
func (j *_JWTToken) DecryptToken(token string) (map[string]interface{}, error) {
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
func (j *_JWTToken) GetTokenContent(token string) (interface{}, error) {
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
func (j *_JWTToken) ParseToken(data string, key []byte) (jwtgo.MapClaims, error) {
	token, err := jwtgo.Parse(data, func(token *jwtgo.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwtgo.MapClaims)

	return claims, nil
}`
