package test

import (
	"amos.wang/awesome/src/main/utils/date_utils"
	"fmt"
	"testing"
	"time"
)

/*
test 文件必须以 _test.go 结尾
功能测试函数 需要以 Test 开头
基准（性能）测试函数 需要以 Benchmark 开头
*/

// 功能测试函数
func TestHello(t *testing.T) {
	str := date_utils.Format(time.Now(), date_utils.Year2Second)
	t.Log("t.Log()", str)

	t.Error(str)
	fmt.Println("after Error")

	t.Fatalf("t.Fatalf >>>>> %v", str)
	// 上边那句执行完,会退出程序,下边这句不会执行
	fmt.Println("after Fatal")
}

// 基准（性能）测试函数
func BenchmarkHello(b *testing.B) {
	date_utils.Format(time.Now(), date_utils.Year2Second)
}
