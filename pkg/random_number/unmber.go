package random_number

import (
	"github.com/spf13/cast"
	"math/rand"
	"time"
)

func GenRandomNumber(l int) string {
	var num string
	// 我们一般使用系统时间的不确定性来进行初始化
	rand.Seed(time.Now().Unix())
	for i := 0; i < l; i++ {
		n := rand.Intn(10)
		num = num + cast.ToString(n)
	}
	return num
}
