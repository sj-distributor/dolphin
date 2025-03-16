package templates

var MiddlewareHandler = `package middleware

import (
	"context"
	"net/http"
	"{{.Config.Package}}/src"
)

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		response.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,SecretKey")
		response.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		response.Header().Set("Access-Control-Expose-Headers", "*")

		ctx := src.NewContext(req.Context())
		ctx.SetHeader(req.Header)

		authorization := ctx.GetHeader("Authorization")
		secretKey := ctx.GetHeader("SecretKey")

		ctx.SetAuthorization(authorization)
		ctx.SetSecretKey(secretKey)

		next.ServeHTTP(response, req.WithContext(ctx.Context()))
	})
}
`
