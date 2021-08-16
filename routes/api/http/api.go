package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *mux.Router) {
	// 加载 admin-api 路由组
	loadAdminApiRouteGroups(router)

	// 加载 api 路由组
	loadApiRouteGroups(router)

	// --- 设定web路由组全局中间件 ---
	//router.Use(middlewares.StartSession)
}

// loadAdminApiRouteGroups admin-api接口路由组
func loadAdminApiRouteGroups(router *mux.Router) {
	adminApiRouter := router.PathPrefix("/admin/api").Subrouter()

	adminApiRouter.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello admin-api, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("admin.api.index")

	// --- 设定 admin-api 路由组全局中间件 ---
	//adminApiRouter.Use()
}

// loadApiRouteGroups api接口路由组
func loadApiRouteGroups(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello api, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("api.index")

	// --- 设定 api 路由组全局中间件 ---
	//adminApiRouter.Use()
}
