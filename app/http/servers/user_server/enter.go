package user_server

import (
	"akita/app/http/dao"
	"sync"
)

type UserServers struct {
	Dao *dao.BaseDao
}

var (
	userServers *UserServers
	once        sync.Once
)

func NewUserServers() *UserServers {
	once.Do(func() {
		if userServers == nil {
			userServers = &UserServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return userServers
}
