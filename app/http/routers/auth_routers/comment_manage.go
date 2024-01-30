package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) CommentManage(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/comment/add	添加评论
	group.POST("comment/add", apis.CommentManage.Add)
	//localhost:8080/api/auth/v1/comment/get	获取评论
	group.GET("comment/get", apis.CommentManage.Get)
}
