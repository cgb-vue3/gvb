package jwt

import (
	"akita/global"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"time"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	ID                   uint   `json:"id"`
	NickName             string `json:"nickname"`
	Role                 int    `json:"role"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(id uint, nickname string, role int) (string, error) {
	// viper获取jwt密钥
	var customSecret = []byte(global.MConfig.JWT.Secret)

	// 创建一个我们自己声明的数据
	claims := CustomClaims{
		ID:       id,
		NickName: nickname,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			// 签发人
			Issuer: "Akita",
		},
	}
	// 生产jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token字符串
	return token.SignedString(customSecret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	// viper获取jwt密钥
	var customSecret = []byte(global.MConfig.JWT.Secret)
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return customSecret, nil
	})

	if err != nil {
		global.Mlog.Error("invalid token", zap.Error(err))
		return nil, err
	}

	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
