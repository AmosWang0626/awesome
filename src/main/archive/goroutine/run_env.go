package main

import (
	"fmt"
	"runtime"
)

func main() {
	nums := runtime.NumCPU()

	// 1.8前设置下以便利用更多CPU
	runtime.GOMAXPROCS(nums - 2)
	fmt.Println("NumCPU", nums)

}
