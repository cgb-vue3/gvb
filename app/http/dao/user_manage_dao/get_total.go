package user_manage_dao

import (
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/get_claims"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M UserManageDao) GetTotal(ctx *gin.Context) {
	// 获取用户总数量
	var (
		userModel    models.UserModel
		userTotal    int64
		articleModel models.ArticleModel
		articleTotal int64
	)

	// 获取用户信息
	id, _, _ := get_claims.GetClaims(ctx)

	// 获取用户总数量
	err := M.Orm.Model(&userModel).Find(&userModel).Count(&userTotal).Error
	if err != nil {
		global.Mlog.Error("获取总数量失败", zap.Error(err))
		response.Err400(ctx)
		return
	}

	// 获取文章总数量
	err = M.Orm.Where("user_model_id = ?", id).Find(&articleModel).Count(&articleTotal).Error
	if err != nil {
		global.Mlog.Error("获取总数量失败", zap.Error(err))
		response.Err400(ctx)
		return
	}

	response.OK200(ctx, response.WithData(map[string]any{
		"userTotal":    userTotal,
		"articleTotal": articleTotal,
	}))
}
