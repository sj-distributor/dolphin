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
func CheckRouterAuth(ctx context.Context, authType int) error {
	if authType == 1 {
		return nil
	}

	var colName string = ""
	resolver := graphql.GetResolverContext(ctx)
	path := utils.StrToArr(resolver.Path().String(), ".")
	if len(path) > 0 {
		colName = path[0]
	}

	if colName == "" {
		return fmt.Errorf("request path is error")
	}

	colName = strcase.ToCamel(colName)

	index := utils.StrIndexOf(NoAuthRoutes, colName)

	if index != -1 {
		return nil
	}

	authorization := ctx.Value("Authorization")
	if authorization == nil {
		return fmt.Errorf("invalid authorization")
	}

	userAgent := ctx.Value("UserAgent")
	if userAgent == nil {
		return fmt.Errorf("unauthorized access")
	}

	// 校验url权限
	err := USER_JWT_TOKEN.Verify(authorization.(string), userAgent.(string))
	return err
}`
