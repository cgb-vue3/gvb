package child_menu_api

import (
	"akita/app/http/controllers/api/v1/auth/child_menu_api/child_menu_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) Add(ctx *gin.Context) {
	var childMenuParams child_menu_params.ChildMenuParams
	err := ctx.ShouldBindJSON(&childMenuParams)
	if err != nil {
		global.Mlog.Error("菜单管理参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.childMenuServers.Add(ctx, childMenuParams)
}
