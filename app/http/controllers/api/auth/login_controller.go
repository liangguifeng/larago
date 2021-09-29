package auth

import (
	"larago/app/http/controllers"
	"larago/pkg/resp"
	"larago/pkg/token"
	"net/http"
)

// LoginController 登录控制器
type LoginController struct {
	controllers.BaseController
}

// Login 首页
func (ac *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: 校验用户输入
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	// 生成token
	generateToken, err := token.GenerateToken(username, password)
	if err != nil {
		resp.JsonResponse(w, 401, "生成token失败，请稍后再试", nil)
		return
	}

	resp.JsonResponse(w, 200, "", map[string]interface{}{
		"data": generateToken,
	})
}
