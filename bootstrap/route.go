package bootstrap

import (
	"github.com/gin-gonic/gin"
	"larago/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()

	//注册Api路由
	routes.RegisterApiRoutes(router)

	//注册Web路由
	routes.RegisterWebRoutes(router)

	return router
}
