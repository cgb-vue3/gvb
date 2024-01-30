package menu_dao

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"akita/app/http/controllers/api/v1/public/menu_api/menu_resp"
	model "akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M MenuDao) Delete(ctx *gin.Context, deleteMenuParams menu_params.DeleteMenuParams, respDeleteMenuInfoList []menu_resp.RespDeleteMenuInfoList) {
	var (
		menuModel []model.MenuModel
		total     int
	)
	// 根据传入的id去第三张表查找判断是否存在关联
	for _, id := range deleteMenuParams.IDList {
		menuRow := M.Orm.Where("id = ?", id).Find(&menuModel).RowsAffected
		// 菜单不存在关联
		if menuRow == 0 {
			global.Mlog.Error("菜单不存在关联")
			respDeleteMenuInfoList = responseList(respDeleteMenuInfoList, id, "菜单不存在关联")
			continue
		}
		// 清除关联
		fmt.Println(menuModel)
		err := M.Orm.Model(&menuModel).Association("ImageModel").Clear()
		if err != nil {
			global.Mlog.Error("菜单关联清除失败")
			respDeleteMenuInfoList = responseList(respDeleteMenuInfoList, id, "菜单关联清除失败")
			continue
		}
		err = M.Orm.Delete(&menuModel).Error
		if err != nil {
			global.Mlog.Error("菜单删除失败", zap.Error(err))
			respDeleteMenuInfoList = responseList(respDeleteMenuInfoList, id, "菜单删除失败")
			continue
		}
		total++
		respDeleteMenuInfoList = responseList(respDeleteMenuInfoList, id, "菜单删除成功")
	}
	response.OK200(
		ctx,
		response.WithData(map[string]any{
			"total":                 total,
			"delete_menu_info_list": respDeleteMenuInfoList,
		}),
	)
}

func responseList(respDeleteMenuInfoList []menu_resp.RespDeleteMenuInfoList, id uint, msg string) []menu_resp.RespDeleteMenuInfoList {
	return append(respDeleteMenuInfoList, menu_resp.RespDeleteMenuInfoList{
		ID:  id,
		Msg: msg,
	})
}
