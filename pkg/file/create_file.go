package file

import (
	"fmt"
	"os"
)

// CreateFile 根据文件路径创建文件
func CreateFile(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
