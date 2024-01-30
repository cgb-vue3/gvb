package response

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeFailed
	CodeCreatedSuccess
	CodeCreatedFailed
	CodeInValidParam
	CodeUserExist
	CodeEmailExist
	CodePhoneExist
	CodeInvalidPassword
	CodePasswordFailed
	CodeUserCreationFailed
	CodeUserCreatedSuccess
	CodeLoginSuccess
	CodeServerBusy
	CodeNoAuth
	CodeNotFound
	CodeValidation
	CodeForbidden
	CodeFileSuffix
	CodeNotImplemented
	CodeInequalityPassword
	CodeGenTokenFailed
	CodeLoginFailed
	CodeUserNoExist
	CodeMenuExists
	CodeMenuNotExists
	CodeMenuCreationFailed
	CodeMenuCreatedSuccess
	CodeSonMenuExists
	CodeSonMenuNotExists
	CodeSonMenuCreationFailed
	CodeSonMenuCreatedSuccess
	CodeGetUserInfoSuccess
	CodeGetUserInfoFailed
	CodeCaptchaError
	CodeChangePwdFailed
	CodeChangePwdSucceed
	CodeWeatherSucceed
	CodeWeatherFailed
	CodeUserListSucceed
	CodeUserListFailed
	CodeImageUploadSucceed
	CodeImageUploadFailed
	CodeImageRespSucceed
	CodeImageRespFailed
	CodeImageRespNil
	CodeUserAddFailed
	CodeUserAddSuccess
	CodeDelSucceed
	CodeDelFailed
	CodeUserUpdateFailed
	CodeUserUpdateSuccess
	CodeReSetPwdFailed
	CodeReSetPwdSuccess
	CodeArticleAddFailed
	CodeArticleAddSuccess
	CodeCommentFailed
	CodeCommentSuccess
	CodeLikeFailed
	CodeLikeSuccess
	CodeArticleGetListFailed
	CodeArticleGetListSuccess
	CodeTagCreatedSucceed
	CodeTagCreatedFailed
	CodeTagRelatedSucceed
	CodeTagRelatedFailed
	CodeTagExist
	CodeArticleDelFailed
	CodeArticleDelSuccess
	CodeGetCommentSuccess
	CodeGetCommentFailed
)

var codeMsgList = map[ResCode]string{
	CodeSuccess:        "成功",
	CodeFailed:         "失败",
	CodeCreatedSuccess: "创建成功",
	CodeCreatedFailed:  "创建失败",
	CodeInValidParam:   "参数错误",
	/*用户*/
	CodeUserNoExist:        "该用户不存在",
	CodeUserExist:          "用户名已存在",
	CodeEmailExist:         "邮箱已存在",
	CodePhoneExist:         "手机号已存在",
	CodeInvalidPassword:    "用户名或密码错误",
	CodePasswordFailed:     "密码错误",
	CodeUserCreationFailed: "注册失败",
	CodeUserCreatedSuccess: "注册成功",
	CodeLoginSuccess:       "登录成功",
	CodeLoginFailed:        "登录失败",
	CodeInequalityPassword: "密码不相等",
	CodeGetUserInfoSuccess: "获取用户信息成功",
	CodeGetUserInfoFailed:  "获取用户信息失败",
	CodeCaptchaError:       "验证码错误或已被使用,请重新发送",
	CodeChangePwdFailed:    "修改密码失败",
	CodeChangePwdSucceed:   "修改密码成功",
	CodeWeatherSucceed:     "天气获取成功",
	CodeWeatherFailed:      "天气获取失败",

	/*用户管理*/
	CodeDelSucceed:        "用户删除成功",
	CodeDelFailed:         "用户删除失败",
	CodeUserAddFailed:     "添加失败",
	CodeUserAddSuccess:    "添加成功",
	CodeUserUpdateFailed:  "编辑失败",
	CodeUserUpdateSuccess: "编辑成功",
	CodeReSetPwdFailed:    "密码重置失败",
	CodeReSetPwdSuccess:   "密码重置成功",

	/*文章管理*/
	CodeArticleAddFailed:      "文章添加失败",
	CodeArticleAddSuccess:     "文章添加成功",
	CodeArticleDelFailed:      "文章删除失败",
	CodeArticleDelSuccess:     "文章删除成功",
	CodeArticleGetListFailed:  "文章列表获取失败",
	CodeArticleGetListSuccess: "文章列表获取成功",
	CodeCommentFailed:         "评论失败",
	CodeCommentSuccess:        "评论成功",
	CodeLikeFailed:            "点赞失败",
	CodeLikeSuccess:           "点赞成功",

	/*文章评论管理*/
	CodeGetCommentSuccess: "获取评论成功",
	CodeGetCommentFailed:  "获取评论失败",
	/*服务器*/
	CodeServerBusy:     "服务繁忙",
	CodeNoAuth:         "权限不足",
	CodeNotFound:       "服务器上没有请求的资源,路径错误等。",
	CodeValidation:     "不允许访问那个资源。该状态码表明对请求资源的访问被服务器拒绝了。（权限，未授权IP等）",
	CodeForbidden:      "服务器拒绝该次访问",
	CodeFileSuffix:     "上传的文件格式不符合规则",
	CodeNotImplemented: "服务器不具备完成请求的功能",
	/*菜单*/
	CodeMenuExists:            "菜单已存在",
	CodeMenuNotExists:         "菜单不存在",
	CodeMenuCreationFailed:    "菜单添加失败",
	CodeMenuCreatedSuccess:    "菜单添加成功",
	CodeSonMenuExists:         "子菜单已存在",
	CodeSonMenuNotExists:      "子菜单不存在",
	CodeSonMenuCreationFailed: "子菜单添加失败",
	CodeSonMenuCreatedSuccess: "子菜单添加成功",
	/*用户管理*/
	CodeUserListSucceed: "用户列表获取成功",
	CodeUserListFailed:  "用户列表获取失败",

	/*标签管理*/
	CodeTagCreatedSucceed: "标签创建成功",
	CodeTagCreatedFailed:  "标签创建失败",
	CodeTagRelatedSucceed: "标签关联成功",
	CodeTagRelatedFailed:  "标签关联失败",
	CodeTagExist:          "标签已存在",
	//CodeTagRelatedFailed:  "标签关联失败",
	/*图片*/
	CodeImageUploadSucceed: "上传图片成功",
	CodeImageUploadFailed:  "上传图片失败",
	CodeImageRespSucceed:   "返回图片成功",
	CodeImageRespFailed:    "返回图片失败",
	CodeImageRespNil:       "暂无图片",
}
