package web

import (
	"larago/app/http/controllers"
	"larago/pkg/view"
	"net/http"
)

// IndexController 首页控制器
type IndexController struct {
	controllers.BaseController
}

// Index 首页
func (ac *IndexController) Index(w http.ResponseWriter, r *http.Request) {
	// ---  2. 加载模板 ---
	view.RenderSimple(w, view.D{}, "welcome.index")
}
