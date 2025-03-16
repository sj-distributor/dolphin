package templates

var ResolverSrc = `package src
import (
	"context"

	"{{.Config.Package}}/gen"
)

func New(db *gen.DB, ec *gen.EventController) gen.Config {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
	// 	return gen.CreateUserHandler(ctx, r, input)
	// }

	// resolver.Handlers.QueryUsers = func(ctx context.Context, r *gen.GeneratedResolver, opts gen.QueryUsersHandlerOptions) (*gen.UserResultType, error) {
	// 	user, err := gen.QueryUserHandler(ctx, r, "userId")
	// 	fmt.Println(user, err)

	// 	return gen.QueryUsersHandler(ctx, r, opts)
	// }

	// events
	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *gen.Event) error {
		return nil
	}

	c := gen.Config{Resolvers: resolver}

	/***
	 * @description: 自定义的 GraphQL 角色校验，用于管理员或者用户对服务的访问权限。
	 *
	 * @param {*string} role - 角色枚举值：ADMIN | USER。
	 *
	 */
	c.Directives.HasRole = func(ctx context.Context, obj any, next graphql.Resolver, role gen.Role) (res any, err error) {
		_, err = next(ctx)

		if err != nil {
			return nil, err
		}

		return value, nil
	}

	/***
	 * @description: 自定义的 GraphQL 验证器函数。
	 *
	 * @param {*string} required - 是否是必填项。如果设置为 "true"，则该字段不能为空。
	 * @param {*string} immutable - 是否是不可修改的。如果设置为 "true"，则不允许修改该字段。
	 * @param {*string} typeArg - 字段类型（例如用于验证的正则表达式），文件路径 utils/rule.go。
	 * @param {*int} minLength - 最小长度，用于字符串长度的验证。
	 * @param {*int} maxLength - 最大长度，用于字符串长度的验证。
	 * @param {*int} minValue - 最小值，用于数值范围的验证。
	 * @param {*int} maxValue - 最大值，用于数值范围的验证。
	 *
	 */
	c.Directives.Validator = func(ctx context.Context, obj any, next graphql.Resolver, required *string, immutable *string, typeArg *string, minLength *int, maxLength *int, minValue *int, maxValue *int) (res any, err error) {
		value, err := next(ctx)

		if err != nil {
			return nil, err
		}

		fieldName := utils.GetFieldName(obj, value)

		if err := utils.ValidateField(ctx, fieldName, value, required, immutable, typeArg, minLength, maxLength, minValue, maxValue); err != nil {
			return nil, err
		}

		// 密码加密
		if typeArg != nil && *typeArg == "password" {
			password := value.(string)
			return utils.EncryptPassword(password), nil
		}

		return value, nil
	}

	return resolver
}
`
