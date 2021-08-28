package resp

import (
	"encoding/json"
	"net/http"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// json 响应封装
func JsonResponse(w http.ResponseWriter, code int, msg string, data interface{}) {
	jsonData, err := json.Marshal(JsonResult{Code: code, Msg: msg,
		Data: data,
	})
	if err != nil {
		// handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, _ = w.Write(jsonData)
}
