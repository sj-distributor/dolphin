package templates

var HttpHandler = `package gen
import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

func GetHTTPServeMux(c Config, db *DB) *mux.Router {
	mux := mux.NewRouter()

	executableSchema := NewExecutableSchema(c)
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
	// 添加上下文数据
	ctx := req.Context()
	ctx = context.WithValue(ctx, KeyLoaders, loaders)
	ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)

	// 返回附带上下文的新请求
	return req.WithContext(ctx)
}
`
