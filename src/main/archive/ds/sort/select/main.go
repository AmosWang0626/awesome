package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 选择排序
// 每次找一个最小的数依次排列
// 第1次: 将最小数放在0位置; 第2次: 将最小数放在1位置; 第3次: 将最小数放在2位置;...
func main() {
	arr := &[...]int{8, 3, 2, 1, 7, 4, 6, 5}
	fmt.Printf("原始数组：%v，数组长度：%d\n", arr, len(arr))
	SelectSorting(arr)

	fmt.Println("\n随机数据执行选择排序")
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
		SelectSorting(arr)
	}

	fmt.Println("Time Consuming", time.Since(t))
}

func SelectSorting(arr *[8]int) {
	sortCount := 0 // 排序次数
	findCount := 0 // 查找次数

	length := len(arr)
	for i := 0; i < length-1; i++ {

		minIndex := i // 最小数坐标
		for j := i; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			findCount++
		}

		//fmt.Printf("[第%d次] 排序前：%v", i+1, arr)
		if minIndex == i {
			//fmt.Println(", 跳过")
			continue
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		//fmt.Printf(", 排序后：%v\n", arr)
		sortCount++
	}

	fmt.Println("排序结果", arr, "\t寻找最小数次数", findCount, "\t排序次数", sortCount)
}
