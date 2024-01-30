package user_params

// RegisterParams 注册参数
type RegisterParams struct {
	NickName   string `json:"nickname" binding:"required,alphaunicode,min=4,max=8"`
	UserName   string `json:"username" binding:"required,alphanum,min=6,max=12"`
	Email      string `json:"email" binding:"required,email"`
	PassWord   string `json:"password" binding:"required,alphanum,min=6,max=12"`
	RePassWord string `json:"re_password" binding:"required,alphanum,min=6,max=12"`
}

// LoginParams 登录参数
type LoginParams struct {
	UserName string `json:"username" binding:"required,alphanum,min=6,max=12"`
	PassWord string `json:"password" binding:"required,alphanum,min=6,max=12"`
}

// EmailParams 通过邮箱发送验证码
type EmailParams struct {
	Email string `json:"email" binding:"required,email"`
}

// ChangePwdParams 修改密码需要的参数
type ChangePwdParams struct {
	UserIDParams
	Code       string `json:"code" binding:"required,max=6"`
	PassWord   string `json:"password" binding:"required,alphanum,min=6,max=12"`
	RePassWord string `json:"re_password" binding:"required,alphanum,min=6,max=12"`
}
type UserIDParams struct {
	ID uint `json:"id" binding:"required"`
}

type AllArticleParams struct {
	ID uint `form:"id"` // 文章id
}
