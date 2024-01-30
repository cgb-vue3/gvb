package upload

import (
	"akita/conf"
	"akita/global"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// DeleteMassQiNiuImages 七牛批量删除图片
func DeleteMassQiNiuImages(fName string) {
	var (
		niu    = global.MConfig.QiNiu
		bucket = niu.BucKet
	)

	manager := bucketManager(niu)
	batch(fName, bucket, manager)

}

func batch(fName, bucket string, manager *storage.BucketManager) {
	var keys []string
	prefix := "images/"
	keys = append(keys, prefix+fName)
	fmt.Println(keys)
	//fmt.Println(keys)
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}
	rets, err := manager.Batch(deleteOps)
	if len(rets) == 0 {
		// 处理错误
		if e, ok := err.(*storage.ErrorInfo); ok {
			global.Mlog.Error(fmt.Sprintf("batch error, code:%v", e.Code))
		} else {
			global.Mlog.Error(fmt.Sprintf("batch error, %s", err))
		}
		return
	}
	// 返回 rets，先判断 rets 是否
	for _, ret := range rets {
		// 200 为成功
		if ret.Code == 200 {
			global.Mlog.Info(fmt.Sprintf("七牛云下%s图片删除成功", prefix+fName))

		} else {
			global.Mlog.Info(fmt.Sprintf("%s\n", ret.Data.Error))
		}
	}
}

func bucketManager(niu *conf.QiNiu) *storage.BucketManager {
	mac := qbox.NewMac(niu.AccessKey, niu.SecretKey)
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(niu.Zone))

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
		Zone:     &zone,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuaBei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager
}
