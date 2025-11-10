/*
 * @Author: Marlon.M
 * @Email: maiguangyang@163.com
 * @Date: 2025-03-26 20:25:34
 */
package templates

var ResolverSrcConfig = `package config

type key int

const (
	// 用户token过期时间
	USER_TOKEN_EXP_TIME = 30

	// 管理员token过期时间
	ADMIN_TOKEN_EXP_TIME = 30

	// 用户token加密key
	USER_TOKEN_SECRET_KEY = "{{.Model.SecretKey}}"

	// 管理员token加密key
	ADMIN_TOKEN_SECRET_KEY = "{{.Model.SecretKey}}"

	KeyHeader        key = iota
	KeyAuthorization key = iota
	KeySecretKey     key = iota
	KeyAppSecret     key = iota
)
`
