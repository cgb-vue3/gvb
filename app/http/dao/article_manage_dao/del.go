package article_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M ArticleManageDao) Del(ctx *gin.Context, delParams article_manage_params.DelParams) {
	var articleModel []models.ArticleModel
	// 通过del参数查找图片
	if err := M.Orm.Where("id in ?", delParams.IDList).Delete(&articleModel).Error; err != nil {
		global.Mlog.Error("文章删除成功", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeArticleDelFailed))
		return
	}
	response.OK200(ctx, response.CodeWithMsg(response.CodeArticleDelSuccess))
}
