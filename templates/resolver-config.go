package templates

var ResolverSrcConfig = `package config
const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "{{.Model.SecretKey}}"
)
`
