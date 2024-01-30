package models

//MessageModel 消息表，用于用户聊天
//type MessageModel struct {
//	gorm.Model
//	SendUserID       uint      `json:"send_user_id"`                           //发送人的ID
//	SendUser         UserModel `json:"send_user" gorm:"foreignKey:SendUserID"` // 发送人
//	SendUserNickName string    `json:"send_user_nick_name"`                    // 发送人的昵称
//	SendUserAvatar   string    `json:"send_user_avatar"`                       // 发送人头像
//	RevUserID        uint      `json:"rev_user_id"`                            //接收人的ID
//	RevUser          UserModel `json:"rev_user" gorm:"foreignKey:RevUserID"`   // 接收人
//	RevUserNickName  string    `json:"rev_user_nick_name"`                     // 接收人的昵称
//	RevUserAvatar    string    `json:"rev_user_avatar"`                        // 接收人头像
//	Content          string    // 消息内容
//}
