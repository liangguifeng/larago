package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// 指定加密密钥
var jwtSecret []byte

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()

	// 默认24小时过期
	expireTime := nowTime.Add(24 * time.Hour)

	// TODO: 生成 token 带上可用信息
	claims := Claims{
		Username: username,
		UserId:   password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer:   "larago",
			IssuedAt: 1630150299,
			Id:       "ssssssssssssss",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
