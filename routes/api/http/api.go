package http

import (
	"github.com/gorilla/mux"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *mux.Router) {
	// 加载 api 路由组
	loadApiRouteGroups(router)

	// --- 设定web路由组全局中间件 ---
	//router.Use()
}

// loadApiRouteGroups api接口路由组
func loadApiRouteGroups(router *mux.Router) {
	//apiRouter := router.PathPrefix("/api").Subrouter()

	//apiRouter.HandleFunc("/index", new(web.IndexController).Index).Methods("GET").Name("admin.api.index")

	// --- 设定 api 路由组全局中间件 ---
	//adminApiRouter.Use()
}
