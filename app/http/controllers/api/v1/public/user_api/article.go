package user_api

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/global"
	"akita/pkg/valida"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (M Api) GetAllArticle(ctx *gin.Context) {
	var all user_params.AllArticleParams
	err := ctx.ShouldBindQuery(&all)
	if err != nil {
		global.Mlog.Error("获取所有文章参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	fmt.Println(all)
	M.userServers.GetAllArticle(ctx, all)
}
