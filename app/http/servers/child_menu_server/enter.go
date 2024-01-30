package child_menu_server

import (
	"akita/app/http/dao"
	"sync"
)

type ChildMenuServers struct {
	Dao *dao.BaseDao
}

var (
	childMenuServers *ChildMenuServers
	once             sync.Once
)

func NewChildMenuServers() *ChildMenuServers {
	once.Do(func() {
		if childMenuServers == nil {
			childMenuServers = &ChildMenuServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return childMenuServers
}
