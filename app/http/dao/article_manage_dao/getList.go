package article_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_resp"
	"akita/app/http/controllers/common"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/paging"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//// respList 获取文章列表返回信息
//type respList struct {
//	ID           uint      `json:"id"`           // 文章id
//	Title        string    `json:"title"`        // 文章标题
//	Abstract     string    `json:"abstract"`     //文章简介
//	Content      string    `json:"content"`      // 文章内容
//	LookCount    int       `json:"lookCount"`    // 文章浏览量
//	CommentCount int       `json:"commentCount"` // 文章评论量
//	LikeCount    int       `json:"likeCount"`    // 文章点赞量
//	Tags         []string  `json:"tags"`         // 文章关联的标签
//	Banner       string    `json:"banner"`       // 文章封面
//	Category     string    `json:"category"`     // 文章分类
//	Issue        time.Time // 发布时间
//}

func (M ArticleManageDao) GetList(ctx *gin.Context, pagParams common.PagingParams, id uint) {
	var (
		articleModel []models.ArticleModel
		list         []article_manage_resp.RespList
		pageSize     = pagParams.PageSize
		offset       int
		sort         string
		count        int64
	)

	// 排序、分页参数
	sort, offset = paging.PagWithSort(pagParams)

	// 通过用户id，查找出对应的文章
	if err := M.Orm.Where("user_model_id = ?", id).Find(&articleModel).Count(&count).Limit(pageSize).Offset(offset).Order(sort).Find(&articleModel).Error; err != nil {
		global.Mlog.Error("文章列表获取失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeArticleGetListFailed))
		return
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
		"total": count,
	}))
}
