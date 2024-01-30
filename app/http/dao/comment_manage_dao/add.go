package comment_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/comment_manage_api/comment_manage_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M CommentManageDao) Add(ctx *gin.Context, commentParams comment_manage_params.CommentParams) {
	//var userModel models.UserModel
	//// 先查询评论者的信息
	//M.Orm.Take(&userModel, "id = ?", commentParams.SelfUserID)
	commentModel := models.CommentModel{
		ArticleModelID: commentParams.ArticleID,
		ParCommentID:   commentParams.ParCommentID,
		Content:        commentParams.Content,
		ParUserID:      commentParams.ParUserID,
		Deep:           commentParams.Deep,
	}
	if err := M.Orm.Create(&commentModel).Error; err != nil {
		global.Mlog.Error("创建评论失败", zap.Error(err))
		response.Err400(ctx)
		return
	}
	response.OK200(ctx)
}
