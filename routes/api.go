package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *mux.Router) {
	// 简单的路由组: api
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("about")
}
