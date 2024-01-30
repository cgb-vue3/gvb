package article_manage_api

import "github.com/gin-gonic/gin"

func (M Api) Put(ctx *gin.Context) {
	M.articleManageServers.Put(ctx)
}
