package user_dao

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/email"
	"akita/pkg/random_number"
	"akita/pkg/redis"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"html/template"
	"time"
)

type RetrievePwd struct {
	SiteName     string
	UserName     string
	UserCode     string
	UserCodeTime int
	SiteAddr     string
}

//type SendLimits struct {
//	count int	// 记录发送次数
//	maxCount int // 最大发送次数
//}

func (M UserDao) SendEmail(ctx *gin.Context, emailParams user_params.EmailParams, temp *template.Template) {
	var (
		userMode   models.UserModel
		expiration = time.Minute * 5 // 过期时间
		html       buffer.Buffer     // 缓存html模板
	)

	// 根据邮箱查找用户判断用户是否存在
	if err := M.Orm.Where("email = ?", emailParams.Email).Take(&userMode).Error; err != nil {
		global.Mlog.Error("该邮箱没有被注册", zap.Error(err))
		response.Err400(ctx, response.WithMsg("该邮箱没有被注册"))
		return
	}
	var id = cast.ToString(userMode.ID)

	// 生产随机验证码，将验证码存入redis并设置过期时间
	num := random_number.GenRandomNumber(6)

	//构造数据渲染模板
	data := RetrievePwd{
		SiteName:     "忘记密码啦，输入验证码重置吧",
		UserName:     userMode.UserName,
		UserCode:     num,
		UserCodeTime: 5,
		SiteAddr:     "https://www.liwenzhou.com/",
	}
	err := temp.Execute(&html, data)
	if err != nil {
		global.Mlog.Error("验证码发送失败", zap.Error(err))
		response.Err400(ctx, response.WithMsg("验证码发送失败"))
		return
	}

	// 将验证码存入redis，key：验证码，value：用户ID
	//rdsData := &RedisData{ID: userMode.ID}
	//marshal, _ := json.Marshal(rdsData)
	set := redis.Set(num, id, expiration)
	if !set {
		global.Mlog.Error("验证码存储失败")
		response.Err400(ctx, response.WithMsg("验证码发送失败"))
		return
	}

	// 发送邮件
	if err := email.InitEmail(emailParams.Email, "重置密码", html.String()); err != nil {
		global.Mlog.Error("验证码发送失败", zap.Error(err))
		response.Err400(ctx, response.WithMsg("验证码发送失败"))
		return
	}

	response.OK200(ctx,
		response.WithMsg("验证码发送成功"),
		response.WithData(map[string]any{
			"id": userMode.ID,
		}),
	)
}
