package valida

import "strings"

// RemoveTopStruct 去掉结构体名称前缀的自定义方法
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
