package templates

var ResolverSrcConfig = `package src

import (
	"{{.Config.Package}}/auth"
)
func Config() {
	auth.USER_JWT_TOKEN.TokenExpTime = 30 // 天数
	auth.USER_JWT_TOKEN.SecretKey = "{{.Model.SecretKey}}"
}`
