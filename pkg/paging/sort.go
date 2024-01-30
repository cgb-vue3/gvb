package paging

import "akita/app/http/controllers/common"

func PagWithSort(pagParams common.PagingParams) (string, int) {
	var sort string
	// 排序
	if pagParams.Sort == 0 {
		// 默认按创建时间排序
		if pagParams.Type == "" {
			sort = "created_at desc"
		} else {
			sort = pagParams.Type + " " + "desc"
		}
	}

	offset := (pagParams.Page - 1) * pagParams.PageSize
	if offset < 0 {
		offset = 0
	}
	return sort, offset
}
