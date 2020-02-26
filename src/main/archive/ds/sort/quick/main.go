package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sortCount = 0 // 插入次数
var swapCount = 0 // 移动数据次数

// 快速排序（Quick Sort）由冒泡排序改进而得
// 1、找一个枢轴(支点)，默认第一个，经过多次交换数据，确定枢轴的位置
// 2、以枢轴的位置，划分左右，所有分别递归（找枢轴，确定枢轴位置，根据枢轴划分左右）
func main() {
	arr := &[...]int{55, 56, 43, 23, 25, 31, 24, 16}
	fmt.Printf("原始数组：%v，数组长度：%d\n", arr, len(arr))
	QuickSorting(arr)

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
		sortCount = 0
		swapCount = 0
		QuickSorting(arr)
	}

	fmt.Println("Time Consuming", time.Since(t))
}

func QuickSorting(arr *[8]int) {
	sort(0, len(arr)-1, arr)
	fmt.Println("排序结果", arr, "\t比较交换次数", swapCount, "\t排序次数", sortCount)
}

// 确定枢轴位置，一分为二，左右分别递归排序
func sort(low, high int, arr *[8]int) {
	if low < high {
		pivotPoint := partition(low, high, arr)
		sort(low, pivotPoint-1, arr)
		sort(pivotPoint+1, high, arr)
	}
}

// 单趟排序，确定枢轴位置
func partition(low, high int, arr *[8]int) int {
	pivotVal := arr[low] // 默认取第一个数
	for low < high {
		for low < high && arr[high] >= pivotVal {
			high--
		}
		arr[low] = arr[high]
		swapCount++
		//fmt.Println("←---比较交换", arr)

		for low < high && arr[low] <= pivotVal {
			low++
		}
		if high == low { // 优化
			break
		}
		arr[high] = arr[low]
		swapCount++
		//fmt.Println("比较交换---→", arr)
	}

	arr[low] = pivotVal
	sortCount++
	//fmt.Println("枢轴位置确定", arr)
	return low
}
