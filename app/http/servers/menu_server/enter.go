package menu_server

import (
	"akita/app/http/dao"
	"sync"
)

type MenuServers struct {
	Dao *dao.BaseDao
}

var (
	menuServers *MenuServers
	once        sync.Once
)

// NewMenuServer 实例化MenuServers
func NewMenuServer() *MenuServers {
	once.Do(func() {
		if menuServers == nil {
			menuServers = &MenuServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return menuServers
}
