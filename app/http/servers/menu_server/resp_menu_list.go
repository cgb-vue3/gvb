package menu_server

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_resp"
	"github.com/gin-gonic/gin"
)

func (M MenuServers) ResponseMenuList(ctx *gin.Context, respMenuList []menu_resp.RespMenuList) {
	M.Dao.MenuDao.ResponseMenuList(ctx, respMenuList)
}
