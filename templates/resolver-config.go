package templates

var ResolverSrcConfig = `package config

type key int

const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "{{.Model.SecretKey}}"

	KeyHeader        key = iota
	KeyAuthorization key = iota
	KeySecretKey     key = iota
)
`
