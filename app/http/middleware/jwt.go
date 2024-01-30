package middleware

import (
	"akita/global"
	"akita/pkg/jwt"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取到请求头中的token
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			global.Mlog.Error("请求头中auth为空")
			response.Err401(ctx, response.WithMsg("请求头中auth为空"))
			ctx.Abort()
			return
		}

		//fmt.Println(authHeader)
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global.Mlog.Error("访问失败,无效的token,请登录!")
			response.Err401(ctx, response.WithMsg("访问失败,无效的token,请登录!"))
			ctx.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			global.Mlog.Error("token解析失败,请重新登录!", zap.Error(err))
			response.Err401(ctx, response.WithMsg("token解析失败,请重新登录!"))
			ctx.Abort()
			return
		}

		//token超时
		if time.Now().Unix() > cast.ToInt64(claims.ExpiresAt.Unix()) {
			global.Mlog.Error("访问失败,token过期,请重新登录!")
			response.Err401(ctx, response.WithMsg("访问失败,token过期,请重新登录!"))
			ctx.Abort() //阻止执行
			return
		}
		// 将当前请求的用户信息保存到请求的上下文c上
		ctx.Set("claims", claims)
		ctx.Next() // 后续的处理函数可以用过c.Get("claims")来获取当前请求的用户信息
	}
}
