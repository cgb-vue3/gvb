package image_server

import (
	"akita/app/http/dao"
	"sync"
)

type ImageServers struct {
	Dao *dao.BaseDao
}

var (
	imageServers *ImageServers
	once         sync.Once
)

// NewImageServers 实例化MenuServers
func NewImageServers() *ImageServers {
	once.Do(func() {
		if imageServers == nil {
			imageServers = &ImageServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return imageServers
}
