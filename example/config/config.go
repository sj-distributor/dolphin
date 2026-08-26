package config

type key int

const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 管理员token过期时间
	ADMIN_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "r8eXaMGD2OrVaDIhzFW0Z9Vm8jjoC0Y0"

	// 管理员token加密key
	ADMIN_TOKEN_SECRET_KEY = "ugMBzU87AQ2iup2kezgdkddC7Rxc6PLE"

	KeyHeader        key = iota
	KeyAuthorization key = iota
	KeySecretKey     key = iota
	KeyAppSecret     key = iota
)
