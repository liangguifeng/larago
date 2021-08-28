package main

import (
	"larago/app/http/middlewares"
	"larago/bootstrap"
	"larago/config"
	c "larago/pkg/config"
	"log"
	"net/http"
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

	log.Println(c.GetString("app.name") + "服务启动成功：http://localhost:" + c.GetString("app.port"))

	// 监听并在 0.0.0.0:3000 上启动服务
	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	if err != nil {
		log.Println(err)
	}
}
