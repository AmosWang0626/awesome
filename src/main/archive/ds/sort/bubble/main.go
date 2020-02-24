package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序（Bubble Sort） 最简单的交换排序
// 每次把最大的数移到最后边，当然，每次都可以少排序一个数
// 第1次: 将最大数放在-1位置; 第2次: 将第二个最大数放在-1位置; 第3次: 将第三个最大数放在-2位置;...
func main() {
	// 特地找了个可以遍历7次的数组
	arr := &[...]int{50, 43, 44, 22, 47, 50, 46, 16}
	fmt.Printf("原始数组：%v，数组长度：%d\n", arr, len(arr))
	sorting(arr)

	fmt.Println("\n随机数据执行冒泡排序")
	testMany()
}

// 测试时
func testMany() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		arr := &[8]int{}
		for j := 0; j < 8; j++ {
			arr[j] = rand.Intn(50) + 10
		}
		fmt.Printf("原始数组：%v\t", arr)
		sorting(arr)
	}
}

func sorting(arr *[8]int) {
	sortCount := 0 // 排序次数
	findCount := 0 // 查找次数

	length := len(arr)
	for i := 0; i < length; i++ {
		temp := 0 // 加个标示, 如果执行下边循环，没有进行过交换，就说明数组有序了
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sortCount++
				temp++
			}
			findCount++
		}
		if temp == 0 {
			fmt.Printf("共需%d次排序, ", i)
			break
		}
	}

	fmt.Println("排序结果", arr, "\t寻找最小数次数", findCount, "\t排序次数", sortCount)
}
