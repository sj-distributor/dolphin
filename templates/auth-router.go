package templates

var AuthRouter = `package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"{{.Config.Package}}/config"
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
	index := utils.StrIndexOf(NoAuthRoutes, methodName)

	if index != -1 {
		return nil
	}

	data, err := UserTokenToMap(ctx)
	if err != nil {
		return err
	}

	content := data["content"].(map[string]interface{})

	if content["role"] == "ADMIN" {
		return AdminTokenVerify(ctx, methodName)
	}

	return UserTokenVerify(ctx, methodName)
}

// 用户token校验
func UserTokenVerify(ctx context.Context, methodName string) error {
	index := utils.StrIndexOf(NoAuthRoutes, methodName)

	if index != -1 {
		return nil
	}

	data, err := UserTokenToMap(ctx)
	if err != nil {
		return err
	}

	content := data["content"].(map[string]interface{})

	// 管理员角色不需要校验
	if content["role"] == "ADMIN" {
		return nil
	}

	token := data["token"].(string)

	return USER_JWT_TOKEN.Verify(token)
}

// 管理员token校验
func AdminTokenVerify(ctx context.Context, methodName string) error {
	index := utils.StrIndexOf(NoAuthRoutes, methodName)

	if index != -1 {
		return nil
	}

	data, err := UserTokenToMap(ctx)
	if err != nil {
		return err
	}

	token := data["token"].(string)

	return ADMIN_JWT_TOKEN.Verify(token)
}
`
