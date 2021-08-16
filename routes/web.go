package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"larago/app/http/middlewares"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *mux.Router) {
	// 加载 admin 路由组
	loadAdminRouteGroups(router)

	// 加载 web 路由组
	loadWebRouteGroups(router)

	// --- 设定web路由组全局中间件 ---
	router.Use(middlewares.StartSession)
}

// adminRouteGroup admin路由组相关路由
func loadAdminRouteGroups(router *mux.Router) {
	// --- 设定路由组前缀 ---
	webRouter := router.PathPrefix("admin").Subrouter()

	// --- 定义路由 ---
	webRouter.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello admin, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("admin.index")

	// --- 设定index路由组中间件 ---
	webRouter.Use(middlewares.StartSession)
}

// webRouteGroup web路由组相关路由
func loadWebRouteGroups(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello web, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("index")

	// --- 设定index路由组中间件 ---
	router.Use(middlewares.StartSession)
}
