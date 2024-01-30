package article_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M ArticleManageDao) Add(ctx *gin.Context, addParams article_manage_params.AddParams) {
	var (
		articleID     uint // 文章id
		tagModel      []models.TagModel
		categoryModel models.CategoryModel
	)
	// 查询分类，取出id
	if err := M.Orm.Take(&categoryModel, "title = ?", addParams.CategoryTitle).Error; err != nil {
		global.Mlog.Error("查询文章分类失败", zap.Error(err))
	}

	//创建文章
	article := &models.ArticleModel{
		UserModelID:     addParams.ID,
		CategoryModelID: categoryModel.ID,
		Title:           addParams.Title,
		Abstract:        addParams.Abstract,
		Content:         addParams.Content,
		Banner:          addParams.Banner,
	}

	if err := M.Orm.Create(article).Error; err != nil {
		global.Mlog.Error("文章添加失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeArticleAddFailed))
		return
	}
	articleID = article.ID
	// 先判断是否有标签，如果有就取出id与文章的id添加到到关联表
	if len(addParams.TagTitles) != 0 {
		if err := M.Orm.Where("title in ?", addParams.TagTitles).Find(&tagModel).Error; err != nil {
			global.Mlog.Error("文章标签查询错误", zap.Error(err))
			return
		}

		// 遍历tagModel，取出id,关联第三张表
		for _, model := range tagModel {
			articleTagModel := models.ArticleTagModel{
				ArticleModelID: articleID,
				TagModelID:     model.ID,
			}

			if err := M.Orm.Create(&articleTagModel).Error; err != nil {
				global.Mlog.Error("文章标签关联失败", zap.Error(err))
				return
			}
		}

	}
	response.OK201(ctx, response.CodeWithMsg(response.CodeArticleAddSuccess))
	return
}
