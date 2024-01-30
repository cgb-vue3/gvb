package exlist

// Existence 判断某个字符是否存在与某个集合中
func Existence(v any, vlist []any) bool {
	for _, value := range vlist {
		if value == v {
			return true
		}
	}
	return false
}
