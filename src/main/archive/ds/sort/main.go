package main

import (
	"amos.wang/awesome/src/main/archive/ds/sort/sort_utils"
	"fmt"
	"math/rand"
	"time"
)

// 快速排序: arr *[80000]int, 耗时: 6.9958ms
// 插入排序: arr *[80000]int, 耗时: 1.1160426s
// 选择排序: arr *[80000]int, 耗时: 4.2486447s
// 冒泡排序: arr *[80000]int, 耗时: 11.3893847s
func main() {
	testMany()
}

func testMany() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		arr := [80000]int{}
		for j := 0; j < 80000; j++ {
			arr[j] = rand.Intn(8000) + 1000
		}

		arr1 := arr
		arr2 := arr
		arr3 := arr
		arr4 := arr

		t := time.Now()
		sort_utils.QuickSorting(&arr1)
		fmt.Printf("快速排序: arr %T, 耗时: %v \n", &arr1, time.Since(t))

		t = time.Now()
		sort_utils.InsertSorting(&arr2)
		fmt.Printf("插入排序: arr %T, 耗时: %v \n", &arr1, time.Since(t))

		t = time.Now()
		sort_utils.SelectSorting(&arr3)
		fmt.Printf("选择排序: arr %T, 耗时: %v \n", &arr1, time.Since(t))

		t = time.Now()
		sort_utils.BubbleSorting(&arr4)
		fmt.Printf("冒泡排序: arr %T, 耗时: %v \n", &arr1, time.Since(t))

		fmt.Println("............................................")
	}

}

// 快速排序: arr *[80000]int, 耗时: 6.9958ms
// 插入排序: arr *[80000]int, 耗时: 1.1160426s
// 选择排序: arr *[80000]int, 耗时: 4.2486447s
// 冒泡排序: arr *[80000]int, 耗时: 11.3893847s
// ............................................
// 快速排序: arr *[80000]int, 耗时: 7.0037ms
// 插入排序: arr *[80000]int, 耗时: 1.2063358s
// 选择排序: arr *[80000]int, 耗时: 4.652509s
// 冒泡排序: arr *[80000]int, 耗时: 12.389444s
// ............................................
// 快速排序: arr *[80000]int, 耗时: 8.9997ms
// 插入排序: arr *[80000]int, 耗时: 1.4000014s
// 选择排序: arr *[80000]int, 耗时: 5.1616514s
// 冒泡排序: arr *[80000]int, 耗时: 11.6379072s
// ............................................
// 快速排序: arr *[80000]int, 耗时: 5.9989ms
// 插入排序: arr *[80000]int, 耗时: 1.1913188s
// 选择排序: arr *[80000]int, 耗时: 4.6810834s
// 冒泡排序: arr *[80000]int, 耗时: 11.774031s
// ............................................
// 快速排序: arr *[80000]int, 耗时: 6.9653ms
// 插入排序: arr *[80000]int, 耗时: 1.2554856s
// 选择排序: arr *[80000]int, 耗时: 4.6385432s
// 冒泡排序: arr *[80000]int, 耗时: 11.4171873s
