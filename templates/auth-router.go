package templates

var AuthRouter = `package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"{{.Config.Package}}/config"
	"{{.Config.Package}}/src/jwt"
	"{{.Config.Package}}/utils"
)

// 获取当前请求的方法名
func GetMethodName(ctx context.Context) (*string, error) {
	var colName string = ""
	resolver := graphql.GetFieldContext(ctx)
	path := utils.StrToArr(resolver.Path().String(), ".")
	if len(path) > 0 {
		colName = path[0]
	}

	if colName == "" {
		return nil, fmt.Errorf("request path is error")
	}

	colName = strcase.ToCamel(colName)

	return &colName, nil
}

// 权限校验
func CheckAuthorization(ctx context.Context, methodName string) (err error) {
	index := utils.StrIndexOf(jwt.NoAuthRoutes, methodName)

	if index != -1 {
		return nil
	}

	authorization := ctx.Value(config.KeyAuthorization)
	if authorization == nil {
		return errors.New("Invalid Authorization")
	}

	token := authorization.(string)

	data, err := parseJWT(token)
	if err != nil {
		return err
	}

	if data["content"] == nil {
		return errors.New("Invalid Authorization")
	}

	content := data["content"].(map[string]interface{})

	if content["role"] == true {
		return AdminTokenVerify(token)
	}

	return UserTokenVerify(token)
}

// 用户token校验
func UserTokenVerify(token string) error {
	// 校验url权限
	err := USER_JWT_TOKEN.Verify(token)
	return err
}

// 管理员token校验
func AdminTokenVerify(token string) error {
	// 校验url权限
	err := ADMIN_JWT_TOKEN.Verify(token)
	return err
}
`
