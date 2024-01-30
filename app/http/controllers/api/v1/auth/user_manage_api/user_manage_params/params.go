package user_manage_params

// AddParams 添加用户参数
type AddParams struct {
	NickName string `json:"nickname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// DelParams 删除用户参数
type DelParams struct {
	IDList []uint `json:"idList"`
}

// PutParams 删除用户参数
type PutParams struct {
	ID       uint   `json:"id" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	PassWord string `json:"password"`
	Type     string `json:"type"`
}
