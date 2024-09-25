package src

import (
	"github.com/sj-distributor/dolphin-example/auth"
)

func Config() {
	auth.USER_JWT_TOKEN.TokenExpTime = 30 // 天数
	auth.USER_JWT_TOKEN.SecretKey = "inXG0fp8Wpe7MRP73mIMaQaxnMhrdHF3"
}
