package main

import "fmt"

// 冒泡排序
func main() {

	arr := []int{45, 13, 24, 57, 18, 23, 77, 26, 33}
	fmt.Println("原始数组", arr)

	bubbleSimple(arr)
	bubblePoint(&arr)

}

func bubbleSimple(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] + arr[j+1]
				arr[j+1] = arr[j] - arr[j+1]
				arr[j] = arr[j] - arr[j+1]
			}
		}
	}
	fmt.Println("冒泡排序后的数组", arr)
}

func bubblePoint(add *[]int) {
	for i := 0; i < len(*add); i++ {
		for j := 0; j < len(*add)-i-1; j++ {
			if (*add)[j] > (*add)[j+1] {
				(*add)[j] = (*add)[j] + (*add)[j+1]
				(*add)[j+1] = (*add)[j] - (*add)[j+1]
				(*add)[j] = (*add)[j] - (*add)[j+1]
			}
		}
	}
	fmt.Println("[指针]冒泡排序后的数组", *add)
}
