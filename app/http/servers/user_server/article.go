package user_server

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"github.com/gin-gonic/gin"
)

func (M UserServers) GetAllArticle(ctx *gin.Context, all user_params.AllArticleParams) {
	M.Dao.UserDao.GetAllArticle(ctx, all)
}
