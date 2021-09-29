package http

import (
	"github.com/gorilla/mux"
	"larago/app/http/controllers/api"
	"larago/app/http/controllers/api/auth"
	"larago/app/http/middlewares"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *mux.Router) {
	// 加载 api 路由组
	loadApiRouteGroups(router)

	// 加载登录路由组
	loadApiAuthRouteGroups(router)

	// 加载需要鉴权的路由组
	loadApiAuthTokenRouteGroups(router)

	// --- 设定 api 路由组全局中间件 ---
	//router.Use()
}

// loadApiRouteGroups api 接口路由组
func loadApiRouteGroups(router *mux.Router) {
	//apiRouter := router.PathPrefix("/api").Subrouter()

	//apiRouter.HandleFunc("/index", new(web.IndexController).Index).Methods("GET").Name("admin.api.index")

	// --- 设定 api 路由组全局中间件 ---
	//adminApiRouter.Use()
}

// loadApiAuthRouteGroups 登录相关路由组
func loadApiAuthRouteGroups(router *mux.Router) {
	// 设定路由组前缀
	apiAuthRouter := router.PathPrefix("/api/auth").Subrouter()

	// 登录路由
	apiAuthRouter.HandleFunc("/login", new(auth.LoginController).Login).Methods("POST").Name("api.auth.login")

	// --- 设定 api-auth 路由组全局中间件 ---
	//apiAuthRouter.Use()
}

// loadApiAuthTokenRouteGroups 需要 token 的路由组
func loadApiAuthTokenRouteGroups(router *mux.Router) {
	// 设定路由组前缀
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/user", new(api.UserController).Index).Methods("GET").Name("api.user.index")

	// --- 验证 token 中间件 ---
	apiRouter.Use(middlewares.VerifyApiToken)
}
