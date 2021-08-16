package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *gin.Engine) {
	// 简单的路由组: v1
	v1 := router.Group("")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
