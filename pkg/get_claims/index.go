package get_claims

import (
	"akita/global"
	"akita/pkg/jwt"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
)

func GetClaims(ctx *gin.Context) (uint, int, string) {
	value, ok := ctx.Get("claims")
	if !ok {
		global.Mlog.Error("权限不足，获取用户信息失败")
		response.Err401(ctx, response.WithMsg("权限不足，获取用户信息失败"))
		return 0, 0, ""
	}
	claims := value.(*jwt.CustomClaims)
	return claims.ID, claims.Role, claims.NickName
}
