package image_dao

import (
	"akita/app/http/models"
)

// Add 添加图片到数据库
func (M ImageDao) Add(name, path, hash string, env models.ImgType) (error, uint) {
	var imgModel = &models.ImageModel{
		Name: name,
		Path: path,
		Hash: hash,
		Env:  env,
	}
	var count int64
	err := M.Orm.Create(&imgModel).Count(&count).Error
	if err != nil {
		return err, 0
	}
	return nil, imgModel.ID
}

// FindImgIsExist 查找图片是否存在
func (M ImageDao) FindImgIsExist(hash string) (bool, error) {
	var count int64
	err := M.Orm.Model(&models.ImageModel{}).Where("hash = ?", hash).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
