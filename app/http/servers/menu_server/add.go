package menu_server

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"github.com/gin-gonic/gin"
)

func (M MenuServers) Add(ctx *gin.Context, menuAddParams menu_params.AddMenuParams) {
	M.Dao.MenuDao.Add(ctx, menuAddParams)
}
