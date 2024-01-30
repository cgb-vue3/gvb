package valida

import (
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validator(ctx *gin.Context, err error) {
	InitTrans()
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.Err422(
			ctx,
			response.CodeWithMsg(response.CodeInValidParam),
			response.WithErr(err),
		)
		return
	}
	response.Err422(
		ctx,
		response.CodeWithMsg(response.CodeInValidParam),
		response.WithValida(RemoveTopStruct(errs.Translate(Trans))),
	)
}
