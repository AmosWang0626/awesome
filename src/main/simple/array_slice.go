package main

import "fmt"

// 观察数组和切片copy区别
func main() {
	/*
		1. 初始化：数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变
		2. 函数传递：数组需要明确指定大小; 切片不需要。数组是值传递，切片是地址传递
	*/

	arr1 := [3]int{1, 2, 3}
	//arr1[3] = 110
	arr2 := arr1
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1, arr2)
	arr2[2] = 1000
	fmt.Println("modified", arr1, arr2)
	printArr(arr3)
	fmt.Println()

	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice3 := make([]int, 5)
	slice4 := make([]int, 5, 10)
	fmt.Println(slice1, slice2)
	slice2[2] = 1000
	fmt.Println("modified", slice1, slice2)
	printSlice(slice3)
	printSlice(slice4)

	printSliceMemoryAddress()
}

func printArr(arr [6]int) {
	fmt.Println(arr)
}

func printSlice(arr []int) {
	fmt.Println(arr)
}

func printSliceMemoryAddress() {
	fmt.Println()
	slice := []int{1, 2, 3, 4}
	// 观察下列元素地址
	for item := range slice {
		fmt.Printf("item=%v, adr=%v, cap=%v\n", item, &item, cap(slice))
	}
	for i := 10; i < 15; i++ {
		slice = append(slice, i)
	}

	fmt.Println()

	// 观察所有元素地址
	for item := range slice {
		fmt.Printf("item=%v, adr=%v, cap=%v\n", item, &item, cap(slice))
	}

}
