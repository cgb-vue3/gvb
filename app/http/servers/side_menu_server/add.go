package side_menu_server

import (
	"akita/app/http/controllers/api/v1/auth/side_menu_api/side_menu_params"
	"github.com/gin-gonic/gin"
)

func (M SideMenuServers) Add(ctx *gin.Context, sideMenuParams side_menu_params.SideMenuParams) {
	M.Dao.SideMenuDao.Add(ctx, sideMenuParams)
}
