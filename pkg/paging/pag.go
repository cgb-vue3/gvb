package paging

import (
	"akita/app/http/controllers/common"
	"gorm.io/gorm"
)

// Option 分页配置，是否开启debug
type Option struct {
	Params common.PagingParams
}

func Pag[T any](model T, Orm *gorm.DB, option Option, choose string) (T, error) {

	var (
		sort string
		//newModel T
		err error
	)

	// 排序
	if option.Params.Sort == 0 {
		sort = "created_at desc"
	}

	offset := (option.Params.Page - 1) * option.Params.PageSize
	if offset < 0 {
		offset = 0
	}
	if choose == "find" {
		// 分页查
		err = Orm.Limit(option.Params.PageSize).Offset(offset).Order(sort).Find(&model).Error

	} else if choose == "model" {
		err = Orm.Model(&model).Limit(option.Params.PageSize).Offset(offset).Order(sort).Error

	}
	return model, err
}
