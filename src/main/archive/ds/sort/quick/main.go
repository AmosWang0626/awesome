package main

import "fmt"

// 快速排序（Quick Sort）由冒泡排序改进而得
// 1、找一个枢轴(支点)，经过多次移动数据，确定枢轴的位置
// 2、以枢轴的位置，划分左右，所有分别递归（找枢轴，确定枢轴位置，根据枢轴划分左右）
func main() {
	arr := &[...]int{8, 3, 2, 1, 7, 4, 6, 5}
	fmt.Printf("原始数组：%v，数组长度：%d\n", arr, len(arr))
	sorting(0, len(arr)-1, arr)
}

// 8, 3, 2, 1, 7, 4, 6, 5
func sorting(left, right int, arr *[8]int) {
	l := left
	r := right
	pivot := arr[(left+right)/2] // 枢轴

	for l < r {
		// 如果想倒叙,很简单,将下边两个for大于小于切换下即可
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	// arr := &[...]int{8, 3, 2, 1, 7, 4, 6, 5}
	// 这个序列, 递归一次就会卡这里, 因为 l == r == 0, 需要错位才能解锁
	if l == r {
		l++
		r--
	}
	// 左递归
	if left < r {
		sorting(left, r, arr)
	}
	// 右递归
	if right > l {
		sorting(l, right, arr)
	}
}
