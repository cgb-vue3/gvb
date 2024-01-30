package category_manage_server

import (
	"akita/app/http/dao"
	"sync"
)

type CategoryServers struct {
	Dao *dao.BaseDao
}

var (
	categoryServers *CategoryServers
	once            sync.Once
)

func NewCategoryServers() *CategoryServers {
	once.Do(func() {
		if categoryServers == nil {
			categoryServers = &CategoryServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return categoryServers
}
