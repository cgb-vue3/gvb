package tag_manage_server

import (
	"akita/app/http/dao"
	"sync"
)

type TagServers struct {
	Dao *dao.BaseDao
}

var (
	tagServers *TagServers
	once       sync.Once
)

func NewTagServers() *TagServers {
	once.Do(func() {
		if tagServers == nil {
			tagServers = &TagServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return tagServers
}
