/*
 * @Author: Marlon.M
 * @Email: maiguangyang@163.com
 * @Date: 2025-11-10 19:08:15
 */
package templates

var AuthUser = `package auth

import (
	"context"
	"errors"

	"{{.Config.Package}}/config"
)

// 用户Token转Map
func UserTokenToMap(ctx context.Context) (content map[string]interface{}, err error) {
	authorization := ctx.Value(config.KeyAuthorization)
	if authorization == nil {
		return content, errors.New("Invalid Authorization")
	}

	token := authorization.(string)

	content, err = ParseJWT(token)

	if err != nil {
		return content, err
	}

	content["token"] = token

	return
}`
