package route

import (
	"github.com/gin-gonic/gin"
)

var route *gin.Engine

// SetRoute 设置路由实例，以供 Name2URL 等函数使用
func SetRoute(r *gin.Engine) {
	route = r
}
