package src

import (
	"github.com/sj-distributor/dolphin-example/auth"
)

func Config() {
	auth.USER_JWT_TOKEN.TokenExpTime = 30 // 天数
	auth.USER_JWT_TOKEN.SecretKey = "6QU5U8oQmWbmywYMZj1eDnVA6QqIST2b"
}
