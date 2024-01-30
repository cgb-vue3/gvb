package public

import (
	"akita/app/http/controllers/api/v1/public/image_api"
	"akita/app/http/controllers/api/v1/public/menu_api"
	"akita/app/http/controllers/api/v1/public/user_api"
)

type PublicGroupApis struct {
	Upload *image_api.Api
	Menu   *menu_api.Api // 菜单
	User   *user_api.Api
}

func NewPublicApis() *PublicGroupApis {
	return &PublicGroupApis{
		image_api.NewImageApi(),
		menu_api.NewMenuApi(),
		user_api.NewUserApi(),
	}
}
