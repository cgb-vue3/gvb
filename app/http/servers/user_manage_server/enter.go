package user_manage_server

import (
	"akita/app/http/dao"
	"sync"
)

type UserManageServers struct {
	Dao *dao.BaseDao
}

var (
	userManageServers *UserManageServers
	once              sync.Once
)

func NewUserManageServers() *UserManageServers {
	once.Do(func() {
		if userManageServers == nil {
			userManageServers = &UserManageServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return userManageServers
}
