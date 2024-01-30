package user_server

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"html/template"
	"os"
)

func (M UserServers) SendEmail(ctx *gin.Context, emailParams user_params.EmailParams) {
	// 解析html文件
	dir, _ := os.Getwd()
	dir = dir + "/templates/email.html"
	temp, err := template.ParseFiles(dir)
	if err != nil {
		global.Mlog.Error("tmpl模板解析错误", zap.Error(err))
		return
	}

	M.Dao.UserDao.SendEmail(ctx, emailParams, temp)
}
