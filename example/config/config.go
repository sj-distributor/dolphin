package config

type key int

const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 管理员token过期时间
	ADMIN_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "3R8RUHm9t3H7GYCHcL8DYoqUVAt2Fh27"

	// 管理员token加密key
	ADMIN_TOKEN_SECRET_KEY = "1WaD8AON28KvSrehsGMCKuqSb1EayzY6"

	KeyHeader        key = iota
	KeyAuthorization key = iota
	KeySecretKey     key = iota
	KeyAppSecret     key = iota
)
