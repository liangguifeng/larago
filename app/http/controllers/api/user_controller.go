package api

import (
	"larago/app/http/controllers"
	"larago/pkg/resp"
	"net/http"
)

// UserController 登录控制器
type UserController struct {
	controllers.BaseController
}

// 获取用户信息
func (ac *UserController) Index(w http.ResponseWriter, r *http.Request) {
	resp.JsonResponse(w, 200, "", map[string]interface{}{
		"username": "admin",
	})
}
