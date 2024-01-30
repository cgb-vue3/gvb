package image_resp

// RespAddImageInfoList 添加图片的参数
type RespAddImageInfoList struct {
	ID        uint   `json:"id,omitempty"`
	FileName  string `json:"file_name,omitempty"` // 文件名
	Url       string `json:"url,omitempty"`       // 图片路径
	IsSuccess bool   `json:"is_success"`          // 是否上传成功
	Size      string `json:"size"`                // 文件大小
	Message   string `json:"message"`             // 提示信息
	//Type      string `json:"type"`                // 图片类型
}

// RespDeleteInfoList 删除的图片信息列表
type RespDeleteInfoList struct {
	ID   uint   `json:"id"`             // 图片的id
	Name string `json:"name,omitempty"` // 图片名
	Msg  string `json:"msg"`            // 删除提示信息
}
