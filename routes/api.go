package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *mux.Router) {
	// 简单的路由组: api
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//可以加上中间件
			log.Println(r.RequestURI)

			next.ServeHTTP(w, r)
		})
	})
	apiRouter.HandleFunc("/ping1", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello1, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("about1")
}
