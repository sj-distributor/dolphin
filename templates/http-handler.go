package templates

var HttpHandler = `package gen
import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"{{.Config.Package}}/auth"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func GetHTTPServeMux(r ResolverRoot, db *DB) *mux.Router {
	mux := mux.NewRouter()
	mux.Use(auth.Handler)

	executableSchema := NewExecutableSchema(Config{Resolvers: r})
	gqlHandler := handler.NewDefaultServer(executableSchema)

	loaders := GetLoaders(db)

	playgroundHandler := HandlerHtml("GraphQL playground", "/graphql")
	mux.HandleFunc("/automigrate", func(res http.ResponseWriter, req *http.Request) {
		err := db.AutoMigrate()
		if err != nil {
			http.Error(res, err.Error(), 400)
		}
		fmt.Fprintf(res, "OK")
	}).Methods("GET")

	// 设置路由
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 使用公共方法处理请求上下文
		r = enrichRequestContext(r, loaders, executableSchema)
		playgroundHandler.ServeHTTP(w, r)
	}).Methods("GET")

	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		// 使用公共方法处理请求上下文
		r = enrichRequestContext(r, loaders, executableSchema)
		gqlHandler.ServeHTTP(w, r)
	}).Methods("POST", "GET")

	handler := mux

	return handler
}

// 公共方法，用于处理请求上下文
func enrichRequestContext(req *http.Request, loaders interface{}, executableSchema interface{}) *http.Request {
	claims, _ := getJWTClaims(req)
	var principalID *string
	if claims != nil {
		if claims["id"] != nil {
			id := claims["id"].(string)
			principalID = &id
		}
	}

	// 添加上下文数据
	ctx := context.WithValue(req.Context(), KeyJWTClaims, claims)
	if principalID != nil {
		ctx = context.WithValue(ctx, KeyPrincipalID, principalID)
	}
	ctx = context.WithValue(ctx, KeyLoaders, loaders)
	ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)

	// 返回附带上下文的新请求
	return req.WithContext(ctx)
}

type JWTClaims struct {
	jwtgo.RegisteredClaims
}

func getJWTClaims(req *http.Request) (res map[string]interface{}, err error) {
	// var p *JWTClaims
	res = map[string]interface{}{}

	tokenStr := strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)

	if tokenStr == "" {
		return
	}

	res, err = auth.USER_JWT_TOKEN.DecryptToken(tokenStr)

	// p = &JWTClaims{}
	// jwtgo.ParseWithClaims(tokenStr, p, nil)
	return res, err
}

var MySecret = []byte("cr6ffSvnPwHwVNgQiQMxtrBtcNRa9NuK")

// 这里传入的是手机号，因为我项目登陆用的是手机号和密码
func MakeToken(phone string) (tokenString string, err error) {
	claim := JWTClaims{
		RegisteredClaims: jwtgo.RegisteredClaims{
			Subject:   phone,
			ExpiresAt: jwtgo.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwtgo.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwtgo.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}
`
