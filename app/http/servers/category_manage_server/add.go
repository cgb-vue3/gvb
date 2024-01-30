package category_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/category_manage_api/category_manage_params"
	"github.com/gin-gonic/gin"
)

func (M CategoryServers) Add(ctx *gin.Context, addParams category_manage_params.AddParams) {
	M.Dao.CategoryManageDao.Add(ctx, addParams)
}
