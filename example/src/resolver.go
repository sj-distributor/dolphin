package src

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sj-distributor/dolphin-example/auth"
	"github.com/sj-distributor/dolphin-example/gen"
	"github.com/sj-distributor/dolphin-example/utils"
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
	 * @param {*string} role - 角色枚举值：ALL | ADMIN | USER。
	 *
	 */
	c.Directives.HasRole = func(ctx context.Context, obj any, next graphql.Resolver, role gen.Role) (res any, err error) {
		methodName, err := auth.GetMethodName(ctx)

		if err != nil {
			return nil, err
		}

		// 验证全部人身份
		if role == gen.RoleAll {

			if err := auth.CheckAuthorization(ctx, *methodName); err != nil {
				return nil, err
			}
		}

		// 验证管理员身份
		if role == gen.RoleAdmin {
			if err := auth.AdminTokenVerify(ctx, *methodName); err != nil {
				return nil, err
			}
		}

		// 验证用户身份
		if role == gen.RoleUser {
			if err := auth.UserTokenVerify(ctx, *methodName); err != nil {
				return nil, err
			}
		}

		return next(ctx)
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
	 * @param {*string} unique - 是否唯一。如果设置为 "true"，则该字段值不能重复。
	 * @param {*string} uniqueScope - 唯一性范围字段（可选）。例如 "uid" 表示在同一 uid 下唯一。
	 *
	 */
	c.Directives.Validator = func(ctx context.Context, obj any, next graphql.Resolver, required *string, immutable *string, typeArg *string, minLength *int, maxLength *int, minValue *int, maxValue *int, unique *string, uniqueScope *string) (res any, err error) {
		value, err := next(ctx)

		if err != nil {
			return nil, err
		}

		fieldName := utils.GetFieldName(obj, value)

		// 注入数据库关联，用于 unique 校验
		if unique != nil && *unique == "true" {
			fc := graphql.GetFieldContext(ctx)
			if fc != nil {
				name := fc.Field.Name
				if strings.HasPrefix(name, "create") {
					name = strings.TrimPrefix(name, "create")
				} else if strings.HasPrefix(name, "update") {
					name = strings.TrimPrefix(name, "update")
				}
				if name != "" && name != fc.Field.Name {
					// 遍历匹配对应的结构体模型，避免直接用字符串转换带来的复数、下划线问题
					var modelStruct any
					for _, v := range gen.TableMap {
						if fmt.Sprintf("%T", v) == "gen."+name {
							modelStruct = v
							break
						}
					}
					if modelStruct != nil {
						ctx = context.WithValue(ctx, "db", db.Query())
						ctx = context.WithValue(ctx, "modelStruct", modelStruct)
					}
				}
			}
		}

		if err := utils.ValidateField(ctx, obj, fieldName, value, required, immutable, typeArg, minLength, maxLength, minValue, maxValue, unique, uniqueScope); err != nil {
			return nil, err
		}

		// 密码加密
		if typeArg != nil && *typeArg == "password" {
			password := value.(string)
			return utils.EncryptPassword(password), nil
		}

		return value, nil
	}

	return c
}
