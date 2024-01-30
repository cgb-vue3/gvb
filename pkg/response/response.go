package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 函数式响应封装

type JSON struct {
	Code          ResCode           `json:"code,omitempty"`
	Message       any               `json:"message,omitempty"`
	Data          map[string]any    `json:"data,omitempty"`
	Error         string            `json:"error,omitempty"`
	ValidaMessage map[string]string `json:"valida_message,omitempty"`
}

func JSON2(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

type RespFNS func(*JSON)

// WithCode 响应code码
func WithCode(code ResCode) RespFNS {
	return func(json *JSON) {
		json.Code = code
	}
}

// WithMsg 响应msg信息
func WithMsg(Message any) RespFNS {
	return func(json *JSON) {
		json.Message = Message
	}
}

// WithData 响应data数据
func WithData(data map[string]any) RespFNS {
	return func(json *JSON) {
		json.Data = data
	}
}

// CodeWithMsg 根据传入的code码，返回对应的code和msg信息
func CodeWithMsg(code ResCode) RespFNS {
	return func(json *JSON) {
		codeMsg, ok := codeMsgList[code]
		if ok {
			json.Code = code
			json.Message = codeMsg
		}
	}
}

// WithErr 响应error
func WithErr(err error) RespFNS {
	return func(json *JSON) {
		json.Error = err.Error()
	}
}

// WithValida 接收参数验证的错误
func WithValida(ValidaErr map[string]string) RespFNS {
	return func(json *JSON) {
		json.ValidaMessage = ValidaErr
	}
}

// <--------------------------------成功得响应封装-------------------------------->

// OK200 成功得通用响应回调，响应状态码200
// 默认返回code码和msg。
//
//	{
//		 "code"：1000
//		 "Message"："您的请求成功了"
//	}
//
// 如果需要返回data数据，在参数中调用WithData即可
// 例：
//
//	response.go.OK(ctx, response.go.WithData(gin.H{
//		"test": "test",
//	}))
func OK200(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeSuccess,
		Message: "您的请求成功了！！",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, defaultJSON)
}

// OK201 响应状态码201
// 请求成功并且服务器创建了新的资源
func OK201(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeSuccess,
		Message: "您的请求成功了！！",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusCreated, defaultJSON)
}

// <--------------------------------失败得响应封装-------------------------------->

// Err400 通用的失败回调，响应状态码 400
// 在解析用户请求，请求的格式或者方法不符合预期时调用
// 默认返回code码和msg。
//
//	{
//		 "code"：1001
//		 "Message"："请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式"
//	}
//
// 如果需要返回data数据，在参数中调用WithData即可
// 例：
//
//	response.go.Err(ctx, response.go.WithData(gin.H{
//		"test": "test",
//	}))
func Err400(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeFailed,
		Message: "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，请确认文件后缀是否符合规则。参数请使用 JSON 格式",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, defaultJSON)
}

// Err401 响应状态码401
// 用户权限不足，未传参时默认调用
// 登录失败，jwt解析失败时请自定义调用
// 如需自定义code和msg请调用WithCode和WithMsg
func Err401(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeNoAuth,
		Message: "权限不足，请确定您有对应的权限",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, defaultJSON)
}

// Err403 响应状态码403
// 服务器拒绝该次访问（如出现跨域）
func Err403(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeForbidden,
		Message: "请求的内容未找到或已删除",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusForbidden, defaultJSON)
}

// Err404 响应状态码404
// 服务器上没有该资源，或者说服务器找不到客户端请求的资源时调用
func Err404(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeNotFound,
		Message: "请求的内容未找到或已删除",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusNotFound, defaultJSON)
}

// Err422 响应状态码422
// 参数错误时响应
func Err422(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeValidation,
		Message: "请求的参数错误，具体请查看 errors",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, defaultJSON)
}

// Err500 服务器错误状态 响应状态码500
func Err500(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeServerBusy,
		Message: "服务器繁忙，请稍后重试",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, defaultJSON)
}

// Err501 服务器错误状态 响应状态码500
func Err501(ctx *gin.Context, fns ...RespFNS) {
	defaultJSON := JSON{
		Code:    CodeServerBusy,
		Message: "尚未实施，或请求格式错误",
	}
	for _, fn := range fns {
		fn(&defaultJSON)
	}
	ctx.AbortWithStatusJSON(http.StatusNotImplemented, defaultJSON)
}
