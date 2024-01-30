package article_manage_server

import "github.com/gin-gonic/gin"

func (M ArticleServers) Put(ctx *gin.Context) {
	M.Dao.ArticleManageDao.Put(ctx)
}
