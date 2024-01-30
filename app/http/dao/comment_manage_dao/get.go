package comment_manage_dao

//
//import (
//	"akita/app/http/controllers/api/v1/auth/comment_manage_api/comment_manage_params"
//	"akita/app/http/models"
//	"akita/global"
//	"akita/pkg/response"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"time"
//)
//
//type ResponseCommentList struct {
//	ArticleID      uint                  `json:"ID"`              // 文章id
//	ParCommentID   uint                  `json:"parID"`           // 父评论id
//	Content        string                `json:"content"`         // 文章内容
//	Author         string                `json:"author"`          // 评论人
//	AuthorAvatar   string                `json:"author_avatar"`   // 作者头像
//	Issue          time.Time             `json:"issue"`           // 评论发布时间
//	SubCommentList []ResponseCommentList `json:"sub_comment"`     // 子评论列表
//	Deep           int                   `json:"deep"`            // 评论深度
//	CommentUserID  uint                  `json:"comment_user_id"` // 评论人id
//	RepUserId      uint                  `json:"rep_user_id"`     // 回复人id
//	Reply          string                `json:"reply"`           // 回复人
//
//}
//
////func (M CommentManageDao) Get(ctx *gin.Context, getParams comment_manage_params.Get) {
////	var (
////		commentModel    []models.CommentModel
////		parCommentModel []ResponseCommentList
////		subCommentModel []ResponseCommentList
////		//ReplyCommentModel []ResponseCommentList
////		respPar = make([]ResponseCommentList, 0)
////	)
////
////	// 查找跟该文章下所有的评论
////	if err := M.Orm.Find(&commentModel, "article_model_id = ?", getParams.ArticleID).Error; err != nil {
////		global.Mlog.Error("查找所有的评论错误", zap.Error(err))
////		response.Err400(ctx, response.CodeWithMsg(response.CodeGetCommentFailed))
////		return
////	}
////	// 遍历评论，将根评论和子评论取出
////	for _, comment := range commentModel {
////		var (
////			par ResponseCommentList
////			sub ResponseCommentList
////			//Rep       ResponseCommentList
////			userModel  models.UserModel
////			RepComment models.UserModel
////		)
////		// 查询出评论用户的信息
////		if err := M.Orm.Where("id = ?", comment.UserModelID).Take(&userModel).Error; err != nil {
////			global.Mlog.Error("查询评论用户错误", zap.Error(err))
////			response.Err400(ctx, response.CodeWithMsg(response.CodeGetCommentFailed))
////			return
////		}
////
////		if comment.Deep == 0 {
////			par.ID = comment.ID
////			par.Content = comment.Content
////			par.ParID = comment.ParCommentID
////			par.Author = userModel.Nickname
////			par.AuthorAvatar = userModel.Avatar
////			par.Issue = comment.CreatedAt
////			par.Deep = comment.Deep
////			par.RepUserId = RepComment.ID
////			par.CommentUserID = userModel.ID
////			//par.Reply = RepComment.Nickname
////			par.SubComment = make([]ResponseCommentList, 0)
////			parCommentModel = append(parCommentModel, par)
////		} else if comment.Deep == 1 {
////			//fmt.Println(comment)
////			// 查出回复用户的信息
////			if err := M.Orm.Where("id = ?", comment.ParUserID).Take(&RepComment).Error; err != nil {
////				global.Mlog.Error("查询评论用户错误", zap.Error(err))
////				response.Err400(ctx, response.CodeWithMsg(response.CodeGetCommentFailed))
////				return
////			}
////			sub.ID = comment.ID
////			sub.Content = comment.Content
////			sub.ParID = comment.ParCommentID
////			sub.Author = userModel.Nickname
////			sub.AuthorAvatar = userModel.Avatar
////			sub.Issue = comment.CreatedAt
////			sub.Deep = comment.Deep
////			sub.RepUserId = RepComment.ID
////			sub.CommentUserID = userModel.ID
////			sub.Reply = RepComment.Nickname
////			sub.SubComment = make([]ResponseCommentList, 0)
////			subCommentModel = append(subCommentModel, sub)
////		}
////	}
////
////	//var data []ResponseCommentList
////	// 循环遍历主、子评论，
////	if len(subCommentModel) != 0 {
////		for _, par := range parCommentModel {
////			var (
////				parList ResponseCommentList
////				subList ResponseCommentList
////				//repList ResponseCommentList
////			)
////			parList = par
////			for _, sub := range subCommentModel {
////				//fmt.Println(sub)
////				fmt.Println(par)
////				if par.ID == sub.ParID {
////					subList = sub
////					parList.SubComment = append(parList.SubComment, subList)
////					//for _, rep := range ReplyCommentModel {
////					//	if sub.ID == rep.ParID {
////					//		rep.Reply = sub.Author
////					//		repList = rep
////					//		sub.Reply = par.Author
////					//		subList = sub
////					//		subList.SubComment = append(subList.SubComment, repList)
////					//		parList.SubComment = append(parList.SubComment, subList)
////					//		continue
////					//	} else {
////					//		sub.Reply = par.Author
////					//		subList = sub
////					//		parList.SubComment = append(parList.SubComment, subList)
////					//
////					//	}
////					//}
////				}
////			}
////			respPar = append(respPar, parList)
////		}
////	}
////
////	response.OK200(ctx, response.WithData(map[string]any{
////		"data": respPar,
////	}), response.CodeWithMsg(response.CodeGetCommentSuccess))
////
////}
//
//func (M CommentManageDao) Get(ctx *gin.Context, getParams comment_manage_params.Get) {
//	var (
//		parCommentModel []models.CommentModel
//		subCommentModel []models.CommentModel
//		//parCommentModel = make([]ResponseCommentList, 0)
//		//subCommentModel []ResponseCommentList
//		//ReplyCommentModel []ResponseCommentList
//		//respPar = make([]ResponseCommentList, 0)
//	)
//
//	// 查找跟该文章下所有的根评论
//	if err := M.Orm.Where("article_model_id = ? AND deep = ?", getParams.ArticleID, 0).Find(&parCommentModel).Error; err != nil {
//		global.Mlog.Error("查找所有的一级评论错误", zap.Error(err))
//		response.Err400(ctx, response.CodeWithMsg(response.CodeGetCommentFailed))
//		return
//	}
//
//	// 查找跟该文章下所有的子评论
//	if err := M.Orm.Where("article_model_id = ? AND deep = ?", getParams.ArticleID, 1).Find(&subCommentModel).Error; err != nil {
//		global.Mlog.Error("查找所有的二级评论错误", zap.Error(err))
//		response.Err400(ctx, response.CodeWithMsg(response.CodeGetCommentFailed))
//		return
//	}
//
//	if len(subCommentModel) > 0 {
//		var (
//			parList ResponseCommentList
//			//subList ResponseCommentList
//			//parList models.CommentModel
//			//subList models.CommentModel
//		)
//
//		// 遍历评论
//		// 对比评论的id与一级评论的父评论id
//		for _, parComment := range parCommentModel {
//			parList.ArticleID = parComment.ArticleModelID
//			parList.ParCommentID = parComment.ParCommentID
//			parList.Issue = parComment.CreatedAt
//			parList.Content = parComment.Content
//		}
//	}
//
//	response.OK200(ctx, response.WithData(map[string]any{
//		"data": subCommentModel,
//	}), response.CodeWithMsg(response.CodeGetCommentSuccess))
//
//}
