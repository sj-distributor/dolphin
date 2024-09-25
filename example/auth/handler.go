package auth

import (
	"context"
	"net/http"
)

type keyType struct {
	name string
}

var UserAgentKey = &keyType{name: "UserAgent"}

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		response.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
		response.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		response.Header().Set("Access-Control-Expose-Headers", "*")

		ctxt := context.WithValue(req.Context(), UserAgentKey, req.Header.Get("User-Agent"))

		if req.Header.Get("Authorization") != "" {
			ctxt = context.WithValue(ctxt, "Authorization", req.Header.Get("Authorization"))
		}

		next.ServeHTTP(response, req.WithContext(ctxt))
	})
}
