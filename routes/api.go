package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes 注册Api相关路由
func RegisterApiRoutes(router *gin.Engine) {
	// 简单的路由组: api
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
