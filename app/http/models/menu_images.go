package models

type MenuImagesModel struct {
	MenuModel    MenuModel  `gorm:"foreignKey:MenuModelID"`
	MenuModelID  uint       `json:"menu_model_id"`
	ImageModel   ImageModel `gorm:"foreignKey:ImageModelID"`
	ImageModelID uint       `json:"image_model_id"`
}
