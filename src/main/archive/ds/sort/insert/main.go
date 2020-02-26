package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 插入排序 内部排序法
// 从arr[i]拿出一个数据val,此时首个要比较的inx为i-1,val小于arr[inx],则inx--继续比较前边的,最后将val赋值给arr[inx+1]
func main() {
	arr := &[...]int{8, 3, 2, 1, 7, 4, 6, 5}
	fmt.Printf("原始数组：%v，数组长度：%d\n", arr, len(arr))
	InsertSorting(arr)

	fmt.Println("\n随机数据执行插入排序")
	testMany()
}

// 测试时
func testMany() {
	t := time.Now()

	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		arr := &[8]int{}
		for j := 0; j < 8; j++ {
			arr[j] = rand.Intn(50) + 10
		}
		fmt.Printf("原始数组：%v\t", arr)
		InsertSorting(arr)
	}

	fmt.Println("Time Consuming", time.Since(t))
}

func InsertSorting(arr *[8]int) {
	sortCount := 0 // 插入次数
	findCount := 0 // 移动数据次数

	length := len(arr)
	for i := 1; i < length; i++ {
		inx := i - 1
		val := arr[i]
		//fmt.Printf("[第%d次] 排序前：%v", i, arr)
		for inx >= 0 && val < arr[inx] {
			arr[inx+1] = arr[inx]
			inx--
			findCount++
		}
		if inx+1 != i {
			arr[inx+1] = val
			sortCount++
		}
		//fmt.Printf(", 排序后：%v\n", arr)
	}

	fmt.Println("排序结果", arr, "\t数据移动次数", findCount, "\t插入次数", sortCount)
}
