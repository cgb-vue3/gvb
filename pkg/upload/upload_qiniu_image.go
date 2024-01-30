package upload

import (
	"akita/conf"
	"akita/global"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

// MassUploadQiNiu 七牛云批量上传图片
func MassUploadQiNiu(data []byte, fileName string) (int, string, string) {
	niu := global.MConfig.QiNiu
	// 判断密钥对是否存在
	if niu.AccessKey == "" || niu.SecretKey == "" {
		return 400, fmt.Sprintf("%s", errors.New("AccessKey或SecretKey未配置")), ""
	}
	// 上传凭证
	upToken := getToken(niu)
	// 配置
	code, msg, url := qiNiuConfig(niu, upToken, data, fileName)
	return code, msg, url
}

// 获取上传凭证
func getToken(niu *conf.QiNiu) string {
	var (
		ak     = niu.AccessKey
		sk     = niu.SecretKey
		bucket = niu.BucKet
	)

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	newMac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(newMac)
	return upToken
}

// 配置参数
func qiNiuConfig(niu *conf.QiNiu, upToken string, data []byte, fileName string) (int, string, string) {
	var imgCdn = global.MConfig.QiNiu.CDN
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(niu.Zone))

	cfg := storage.Config{
		Zone:          &zone, // 地区
		UseHTTPS:      false, //非https
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&cfg) // 上传后返回的结果

	ret := storage.PutRet{} // 上传后返回的结果

	putExtra := storage.PutExtra{ // 额外参数
		Params: map[string]string{},
	}
	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := "images/" + fileName // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), int64(len(data)), &putExtra)
	if err != nil {
		global.Mlog.Error("文件上传七牛云失败：%s", zap.Error(err))
		return 501, err.Error(), ""
	}
	url := imgCdn + ret.Key // 返回上传后的文件访问路径
	return 201, "上传成功", url
}
