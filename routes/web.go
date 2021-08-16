package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(router *mux.Router) {
	// 简单的路由组: v1
	router.HandleFunc("/api/ping1", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	}).Methods("GET").Name("about")
}
