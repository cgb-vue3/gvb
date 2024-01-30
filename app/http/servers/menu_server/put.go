package menu_server

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"github.com/gin-gonic/gin"
)

func (M MenuServers) Put(ctx *gin.Context, putMenuParam menu_params.PutMenuParams) {
	M.Dao.MenuDao.Put(ctx, putMenuParam)
}
