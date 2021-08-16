package bootstrap

import (
	"github.com/gorilla/mux"
	"larago/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()

	//注册Api路由
	routes.RegisterApiRoutes(router)

	//注册Web路由
	routes.RegisterWebRoutes(router)

	return router
}
