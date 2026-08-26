package auth

import (
	"context"
	"errors"

	"github.com/sj-distributor/dolphin-example/config"
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
}
