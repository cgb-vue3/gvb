package uuid

import (
	"github.com/google/uuid"
	"strings"
)

// GetStrUUID 生成uuid字符串
func GetStrUUID(l int) string {
	random, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	s := random.String()
	rs := strings.ReplaceAll(s, "-", "")

	if l > 32 {
		l = 32
	}
	return rs[:l]
}
