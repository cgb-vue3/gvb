package comment_manage_server

import (
	"akita/app/http/dao"
	"sync"
)

type CommentServers struct {
	Dao *dao.BaseDao
}

var (
	commentServers *CommentServers
	once           sync.Once
)

func NewCommentServers() *CommentServers {
	once.Do(func() {
		if commentServers == nil {
			commentServers = &CommentServers{
				Dao: dao.NewBaseDao(),
			}
		}
	})
	return commentServers
}
