package templates

var AuthJWT = `package auth

import (
	"errors"
	"time"

	jwtgo "github.com/golang-jwt/jwt/v5"
	"{{.Config.Package}}/config"
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
	jwt.RegisteredClaims
}

type JWTToken struct {
	TokenExpTime int64
	SecretKey    string
}

// 設置JWT
func (j *JWTToken) SetToken(str interface{}) (string, error) {
	timeNow := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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
func (j *JWTToken) ParseToken(data string, key []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

// 解析 JWT 中间部分（Payload）
func ParseJWT(token string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	payloadSegment := parts[1]

	// Base64 解码
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %v", err)
	}

	// JSON 解析
	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return payload, nil
}
`
