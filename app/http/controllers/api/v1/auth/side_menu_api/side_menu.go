package side_menu_api

import (
	"akita/app/http/controllers/api/v1/auth/side_menu_api/side_menu_params"
	"akita/app/http/controllers/common"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

// Add 添加菜单
func (M Api) Add(ctx *gin.Context) {
	var sideMenuParams side_menu_params.SideMenuParams
	err := ctx.ShouldBindJSON(&sideMenuParams)
	if err != nil {
		global.Mlog.Error("菜单管理参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.sideMenuServers.Add(ctx, sideMenuParams)
}

// GetPag 获取菜单
func (M Api) GetPag(ctx *gin.Context) {
	var (
		sidePagParams common.PagingParams
	)
	err := ctx.ShouldBindJSON(&sidePagParams)
	if err != nil {
		global.Mlog.Error("菜单管理参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.sideMenuServers.GetPag(ctx, sidePagParams)
}
