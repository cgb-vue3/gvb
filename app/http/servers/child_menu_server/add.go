package child_menu_server

import (
	"akita/app/http/controllers/api/v1/auth/child_menu_api/child_menu_params"
	"github.com/gin-gonic/gin"
)

func (M ChildMenuServers) Add(ctx *gin.Context, childMenuParams child_menu_params.ChildMenuParams) {
	M.Dao.ChildMenuDao.Add(ctx, childMenuParams)
}
