package middlewares

import (
	"fmt"
	"larago/pkg/resp"
	"larago/pkg/token"
	"net/http"
	"strings"
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

		// 开始校验token
		expiresAt := parseToken.StandardClaims.ExpiresAt
		fmt.Println(expiresAt)
		//if expiresAt > time.Now() {
		//
		//}

		// 将请求传递下去
		next.ServeHTTP(w, r)
	})
}
