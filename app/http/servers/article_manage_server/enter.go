package article_manage_server

import (
	"akita/app/http/dao"
	"sync"
)

type ArticleServers struct {
	Dao *dao.BaseDao
}

var (
	articleServers *ArticleServers
	once           sync.Once
)

func NewArticleServers() *ArticleServers {
	once.Do(func() {
		if articleServers == nil {
			articleServers = &ArticleServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return articleServers
}
