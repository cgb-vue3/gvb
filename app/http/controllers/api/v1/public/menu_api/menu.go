package menu_api

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"akita/app/http/controllers/api/v1/public/menu_api/menu_resp"
	"akita/global"
	"akita/pkg/valida"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Add 添加菜单
func (M Api) Add(ctx *gin.Context) {
	var menuAddParams menu_params.AddMenuParams
	err := ctx.ShouldBindJSON(&menuAddParams)
	if err != nil {
		global.Mlog.Error("参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.menuServer.Add(ctx, menuAddParams)
}

// ResponseMenuList 返回菜单列表
func (M Api) ResponseMenuList(ctx *gin.Context) {
	var respMenuList = make([]menu_resp.RespMenuList, 0)
	M.menuServer.ResponseMenuList(ctx, respMenuList)
}

// Put 更新菜单
func (M Api) Put(ctx *gin.Context) {
	var putMenuParam menu_params.PutMenuParams
	fmt.Println(putMenuParam)
	err := ctx.ShouldBindJSON(&putMenuParam)
	if err != nil {
		valida.Validator(ctx, err)
		return
	}
	M.menuServer.Put(ctx, putMenuParam)
}

// Delete 删除菜单
func (M Api) Delete(ctx *gin.Context) {
	var deleteMenuParams menu_params.DeleteMenuParams
	var respDeleteMenuInfoList = make([]menu_resp.RespDeleteMenuInfoList, 0)
	err := ctx.ShouldBindJSON(&deleteMenuParams)
	if err != nil {
		valida.Validator(ctx, err)
		return
	}
	M.menuServer.Delete(ctx, deleteMenuParams, respDeleteMenuInfoList)
}
