package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) ArticleManage(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/article/add	添加文章
	group.POST("article/add", apis.ArticleManage.Add)
	//localhost:8080/api/auth/v1/article/getList	返回文章列表
	group.GET("article/getList", apis.ArticleManage.GetList)
	//localhost:8080/api/auth/v1/article/put	编辑文章
	group.PUT("article/put", apis.ArticleManage.Put)
	//localhost:8080/api/auth/v1/article/del	删除文章
	group.DELETE("article/del", apis.ArticleManage.Del)
}
