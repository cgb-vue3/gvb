package image_api

import "akita/app/http/servers/Image_Server"

type Api struct {
	imageServer *image_server.ImageServers // 图片
}

func NewImageApi() *Api {
	return &Api{
		imageServer: image_server.NewImageServers(),
	}
}
