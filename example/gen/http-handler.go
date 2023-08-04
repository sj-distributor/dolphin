package gen

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/sj-distributor/dolphin-example/auth"

	jwtgo "github.com/golang-jwt/jwt/v5"
)

func GetHTTPServeMux(r ResolverRoot, db *DB) *mux.Router {
	mux := mux.NewRouter()
	mux.Use(auth.Handler)

	executableSchema := NewExecutableSchema(Config{Resolvers: r})
	gqlHandler := handler.NewDefaultServer(executableSchema)

	loaders := GetLoaders(db)

	// playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
	playgroundHandler := HandlerHtml("GraphQL playground", "/graphql")
	mux.HandleFunc("/automigrate", func(res http.ResponseWriter, req *http.Request) {
		err := db.AutoMigrate()
		if err != nil {
			http.Error(res, err.Error(), 400)
		}
		fmt.Fprintf(res, "OK")
	})
	mux.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		claims, _ := getJWTClaims(req)
		var principalID *string
		if claims != nil {
			principalID = &(*claims).Subject
		}
		ctx := context.WithValue(req.Context(), KeyJWTClaims, claims)
		if principalID != nil {
			ctx = context.WithValue(ctx, KeyPrincipalID, principalID)
		}
		ctx = context.WithValue(ctx, KeyLoaders, loaders)
		ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
		req = req.WithContext(ctx)
		if req.Method == "GET" {
			playgroundHandler(res, req)
		} else {
			gqlHandler.ServeHTTP(res, req)
		}
	})
	handler := mux

	return handler
}

type JWTClaims struct {
	jwtgo.RegisteredClaims
}

func getJWTClaims(req *http.Request) (*JWTClaims, error) {
	var p *JWTClaims

	tokenStr := strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)
	if tokenStr == "" {
		return p, nil
	}

	p = &JWTClaims{}
	jwtgo.ParseWithClaims(tokenStr, p, nil)
	return p, nil
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
