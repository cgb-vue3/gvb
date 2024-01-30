package image_server

import (
	"akita/app/http/controllers/api/v1/public/image_api/image_resp"
	"akita/global"
	"akita/pkg/encryption"
	"akita/pkg/file"
	"akita/pkg/response"
	"akita/pkg/upload"
	"akita/pkg/uuid"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"path"
)

// Add 处理上传图片相关的请求
func (M ImageServers) Add(ctx *gin.Context, respAddImageInfoList []image_resp.RespAddImageInfoList) {
	var total int // 统计添加成功的条数
	getPath := global.MConfig.Upload.FilePath
	maxsize := global.MConfig.Upload.FileSize

	isCreateFile(getPath)

	form, err := ctx.MultipartForm()
	if err != nil {
		response.Err400(ctx, response.CodeWithMsg(response.CodeImageUploadFailed))
		return
	}

	files := form.File["images"]
	respAddImageInfoList = M.processImages(ctx, files, getPath, maxsize, respAddImageInfoList)
	for _, imgParams := range respAddImageInfoList {
		if imgParams.IsSuccess {
			total++
		}
	}

	response.OK200(
		ctx,
		response.CodeWithMsg(response.CodeImageUploadSucceed),
		response.WithData(map[string]any{
			"total":         total,
			"add_file_list": respAddImageInfoList,
		}))
}

// 对图片进行加工
func (M ImageServers) processImages(ctx *gin.Context, files []*multipart.FileHeader, getPath string, maxsize float64, respAddImageInfoList []image_resp.RespAddImageInfoList) []image_resp.RespAddImageInfoList {
	for _, image := range files {
		open, err := image.Open()
		if err != nil {
			continue
		}
		readData, err := io.ReadAll(open)
		if err != nil {
			continue
		}

		// 获取随机数，替换掉文件名
		rdName := uuid.GetStrUUID(10)
		// 获取后缀
		ext := path.Ext(image.Filename)
		fName := rdName + ext
		// 拼接图片路径
		fPath := path.Join(getPath, fName)
		// 计算图片大小，单位MB
		fSize := float64(image.Size) / float64(1024*1024)

		// 白名单验证
		if !upload.White(ext) {
			respAddImageInfoList = addImageList(0, fName, "", "上传文件的后缀不符合规则", false, fSize, respAddImageInfoList)
			continue
		}
		// 文件大小验证
		if fSize >= maxsize {
			respAddImageInfoList = addImageList(0, fName, "", "上传文件过大,最大为"+cast.ToString(maxsize)+"MB", false, fSize, respAddImageInfoList)
			continue
		}

		// MD5生成加密字符串
		md5Str := encryption.MD5(readData)

		// 查图片在数据库中是否存在
		exist, _ := M.Dao.FindImgIsExist(md5Str)
		if exist {
			respAddImageInfoList = addImageList(0, fName, "", "图片已存在", false, fSize, respAddImageInfoList)
			continue
		}

		// 是否上传七牛云
		if global.MConfig.QiNiu.Enable {
			// 七牛云批量上传图片
			code, msg, url := upload.MassUploadQiNiu(readData, fName)
			if code == 400 {
				respAddImageInfoList = addImageList(0, fName, url, msg, false, fSize, respAddImageInfoList)
				continue
			} else if code == 501 {
				respAddImageInfoList = addImageList(0, fName, url, msg, false, fSize, respAddImageInfoList)
			} else {
				// 图片入库
				err, id := M.Dao.ImageDao.Add(fName, url, md5Str, 1)
				if err != nil {
					return nil
				}
				respAddImageInfoList = addImageList(id, fName, url, msg, true, fSize, respAddImageInfoList)
				continue
			}
		}

		// 保存到本地
		err = ctx.SaveUploadedFile(image, fPath)
		if err != nil {
			global.Mlog.Error("本地保存文件失败：%s", zap.Error(err))
			respAddImageInfoList = addImageList(0, fName, fPath, "文件上传失败", false, fSize, respAddImageInfoList)
			continue
		}
		// 图片入库
		err, id := M.Dao.ImageDao.Add(fName, fPath, md5Str, 0)
		if err != nil {
			return nil
		}
		respAddImageInfoList = addImageList(id, fName, fPath, "文件成功上传到本地", true, fSize, respAddImageInfoList)
	}

	return respAddImageInfoList
}

// 判断本地文件夹是否存在
func isCreateFile(getPath string) {
	if getPath == "" {
		getPath = "uploads/images"
	}

	exists, _ := file.PathExists(getPath)
	if !exists {
		global.Mlog.Warn("文件夹不存在，将自动创建")
		file.CreateFile(getPath)
		global.Mlog.Info(fmt.Sprintf("文件夹创建成功：%s", getPath))
	}
}

// 追加图片信息
func addImageList(id uint, fName, url, msg string, success bool, imgSize float64, fileList []image_resp.RespAddImageInfoList) []image_resp.RespAddImageInfoList {
	return append(fileList, image_resp.RespAddImageInfoList{
		ID:        id,
		FileName:  fName,
		Url:       url,
		IsSuccess: success,
		Size:      fmt.Sprintf("%.2fMB", imgSize),
		Message:   msg,
	})
}
