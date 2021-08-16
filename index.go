package main

import (
	"larago/bootstrap"
	"larago/config"
	c "larago/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	// 监听并在 0.0.0.0:8080 上启动服务
	_ = router.Run(":" + c.GetString("app.port"))
}
