package menu_server

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"akita/app/http/controllers/api/v1/public/menu_api/menu_resp"
	"github.com/gin-gonic/gin"
)

func (M MenuServers) Delete(ctx *gin.Context, deleteMenuParams menu_params.DeleteMenuParams, respDeleteMenuInfoList []menu_resp.RespDeleteMenuInfoList) {
	M.Dao.MenuDao.Delete(ctx, deleteMenuParams, respDeleteMenuInfoList)
}
