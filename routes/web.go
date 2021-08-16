package routes

import (
	"github.com/gorilla/mux"
	"larago/app/http/controllers/web"
	"larago/app/http/middlewares"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *mux.Router) {
	// 加载 web 路由组
	loadWebRouteGroups(router)
	// 加载 静态资源 路由组
	loadStaticRouteGroups(router)

	// --- 设定web路由组全局中间件 ---
	//router.Use()
}

// webRouteGroup web路由组相关路由
func loadWebRouteGroups(router *mux.Router) {
	router.HandleFunc("/", new(web.IndexController).Index).Methods("GET").Name("index")

	// --- 设定index路由组中间件 ---
	router.Use(middlewares.StartSession)
}

// loadStaticRouteGroups 静态资源路由组相关路由
func loadStaticRouteGroups(router *mux.Router) {
	// 静态资源
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./resources/static")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./resources/static")))
}
