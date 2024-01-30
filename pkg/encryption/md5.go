package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 数字加密
func MD5(data []byte) string {
	sum := md5.Sum(data)
	toString := hex.EncodeToString(sum[:])
	return toString
}
