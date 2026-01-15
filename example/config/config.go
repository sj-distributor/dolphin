package config

type key int

const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 管理员token过期时间
	ADMIN_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "Tz3gfYg41DyNwr0blefX6DgFpoiD7tze"

	// 管理员token加密key
	ADMIN_TOKEN_SECRET_KEY = "9hH2p6QGhUH5y4XwTXDaQo5DRNFIFtKf"

	KeyHeader        key = iota
	KeyAuthorization key = iota
	KeySecretKey     key = iota
	KeyAppSecret     key = iota
)
