package middlewares

import (
	"larago/pkg/resp"
	"larago/pkg/token"
	"net/http"
	"strings"
	"time"
)

// VerifyApiToken
func VerifyApiToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取token
		authorization := r.Header.Get("Authorization")

		// 拆分token
		tokenString := strings.Fields(authorization)

		// 解析token
		parseToken, err := token.ParseToken(tokenString[1])
		if err != nil {
			resp.JsonResponse(w, 401, "token校验失败，请重新获取", nil)
			return
		}

		// 获取 token 到期时间
		expiresAt := parseToken.StandardClaims.ExpiresAt

		// 开始校验 token
		// 1. 校验 token 是否过期
		if time.Now().Unix() > expiresAt {
			resp.JsonResponse(w, 401, "token已过期，请重新获取", nil)
			return
		}

		// 将 jwt 信息放入请求头，方便控制器获取
		r.Header.Set("userId", parseToken.UserId)
		r.Header.Set("userName", parseToken.Username)
		r.Header.Set("projectId", parseToken.StandardClaims.Id)

		// 将请求传递下去
		next.ServeHTTP(w, r)
	})
}
