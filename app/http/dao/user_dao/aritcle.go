package user_dao

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_resp"
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func (M UserDao) GetAllArticle(ctx *gin.Context, all user_params.AllArticleParams) {
	//var articleModel models.ArticleTagModel

	var (
		articleModel []models.ArticleModel
		list         []article_manage_resp.RespList
	)

	if all.ID != cast.ToUint(nil) {
		M.Orm.Where("id = ?", all.ID).Find(&articleModel)
	} else {
		M.Orm.Find(&articleModel)
	}

	// 循环遍历文章，通过文章id去查每个文章的分类和标签
	for _, model := range articleModel {
		var (
			category   models.CategoryModel
			tag        []models.TagModel
			userModel  models.UserModel
			articleTag []models.ArticleTagModel
			tagTitle   = make([]string, 0)
		)
		// 查询文章作者
		if err := M.Orm.Where("id = ?", model.UserModelID).Find(&userModel).Error; err != nil {
			global.Mlog.Error("文章作者查询失败", zap.Error(err))
			return
		}

		// 查询文章的分类
		if err := M.Orm.Where("id = ?", model.CategoryModelID).Find(&category).Error; err != nil {
			global.Mlog.Error("文章分类查询失败", zap.Error(err))
			return
		}
		// 查询关联表的id
		if err := M.Orm.Where("article_model_id = ?", model.ID).Find(&articleTag).Error; err != nil {
			global.Mlog.Error("关联表查询失败", zap.Error(err))
			return
		}

		// 循环遍历关联表，取出标签id
		for _, articleTagModel := range articleTag {
			// 查询文章的标签
			if err := M.Orm.Where("id = ?", articleTagModel.TagModelID).Find(&tag).Error; err != nil {
				global.Mlog.Error("文章标签查询失败", zap.Error(err))
				return
			}
			for _, tagModel := range tag {
				tagTitle = append(tagTitle, tagModel.Title)
			}
		}

		var articleData = article_manage_resp.RespList{
			ID:           model.ID,
			Author:       userModel.Nickname,
			AuthorAvatar: userModel.Avatar,
			Title:        model.Title,
			Abstract:     model.Abstract,
			Content:      model.Content,
			LookCount:    model.LookCount,
			CommentCount: model.CommentCount,
			LikeCount:    model.LikeCount,
			Tags:         tagTitle,
			Banner:       model.Banner,
			Category:     category.Title,
			Issue:        model.CreatedAt,
		}
		list = append(list, articleData)
	}

	response.OK200(ctx, response.CodeWithMsg(response.CodeArticleGetListSuccess), response.WithData(map[string]any{
		"data":  list,
		"total": len(list),
	}))
}
