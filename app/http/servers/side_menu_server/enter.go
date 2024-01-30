package side_menu_server

import (
	"akita/app/http/dao"
	"sync"
)

type SideMenuServers struct {
	Dao *dao.BaseDao
}

var (
	sideMenuServers *SideMenuServers
	once            sync.Once
)

func NewSideMenuServers() *SideMenuServers {
	once.Do(func() {
		if sideMenuServers == nil {
			sideMenuServers = &SideMenuServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return sideMenuServers
}
