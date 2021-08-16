package bootstrap

import (
	"github.com/gorilla/mux"
	"larago/routes"
	"larago/routes/api/http"
	"larago/routes/api/rpc"
	"larago/routes/api/socket"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()

	//注册http-api路由
	http.RegisterApiRoutes(router)

	//注册rpc-api路由
	rpc.RegisterApiRoutes(router)

	//注册socket-api路由
	socket.RegisterApiRoutes(router)

	//注册Web路由
	routes.RegisterWebRoutes(router)

	return router
}
