package templates

var AuthRouter = `package auth

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iancoleman/strcase"
	"{{.Config.Package}}/utils"
)

// 检测路由是否需要登录
func CheckRouterAuth(ctx context.Context) error {
	var colName string = ""
	resolver := graphql.GetFieldContext(ctx)
	path := utils.StrToArr(resolver.Path().String(), ".")
	if len(path) > 0 {
		colName = path[0]
	}

	if colName == "" {
		return fmt.Errorf("request path is error")
	}

	colName = strcase.ToCamel(colName)

	err := CheckAuthorization(ctx, colName)
	return err
}

// CheckAuthorization ....
func CheckAuthorization(ctx context.Context, colName string) error {
	index := utils.StrIndexOf(NoAuthRoutes, colName)

	if index != -1 {
		return nil
	}

	authorization := ctx.Value("Authorization")
	if authorization == nil {
		return errors.New("Invalid Authorization")
	}

	// 校验url权限
	err := USER_JWT_TOKEN.Verify(authorization.(string))
	return err
}
`
