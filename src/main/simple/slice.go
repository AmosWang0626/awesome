package main

import "fmt"

func main() {
	nums := []int{100, 200, 300, 400, 500, 600, 700, 800}
	create(nums)
	query(nums)
	myAppend(nums)
	myCopy()
	fmt.Println("..............................")
	str := "hello#zzti.com"
	fmt.Println(str)
	fmt.Println("str[:5]", str[:5])
	// 英文
	arr1 := []byte(str)
	arr1[5] = '@'
	str = string(arr1)
	fmt.Println(str)
	// 中文
	arr2 := []rune(str)
	arr2[5] = '哈'
	str = string(arr2)
	fmt.Println(str)
	fmt.Println("..............................")

}

func myAppend(nums []int) {
	myNew := append(nums, 900, 1000)
	fmt.Println("\nmyNew =", myNew, ", len =", len(myNew), ", cap =", cap(myNew))
	myNew = append(nums, myNew...)
	myNew = append(nums, nums...)
	fmt.Println("\nmyNew =", myNew, ", len =", len(myNew), ", cap =", cap(myNew))
}

func myCopy() {
	var slice1 = []int{100, 200, 300}
	var slice2 = make([]int, 8)
	fmt.Println("\nslice1 =", slice1, ", len =", len(slice1), ", cap =", cap(slice1))
	// copy(target, src)
	copy(slice2, slice1)
	fmt.Println("do copy\nslice2 =", slice2, ", len =", len(slice2), ", cap =", cap(slice2))

	var slice3 = []int{100, 200, 300}
	var slice4 = make([]int, 1)
	fmt.Println("\nslice4 =", slice4, ", len =", len(slice4), ", cap =", cap(slice4))
	copy(slice4, slice3)
	fmt.Println("do copy\nslice4 =", slice4, ", len =", len(slice4), ", cap =", cap(slice4))

}

func create(nums []int) {
	// 1.1 创建切片方式1
	// 数组是可见的, 可通过数组和切片进行访问
	slice1 := nums[1:3]
	// 含头不含尾
	fmt.Println("slice1 =", slice1, ", len =", len(slice1), ", cap =", cap(slice1))

	// 1.2 创建切片方式2
	// 使用内置函数 make(type, len, cap)
	// 数组是不可见的, 只能通过切片进行访问
	slice2 := make([]int, 4, 8)
	fmt.Println("slice2 =", slice2, ", len =", len(slice2), ", cap =", cap(slice2))
	slice2[0] = 100
	slice2[2] = 200
	fmt.Println("\tchanged =", slice2, ", len =", len(slice2), ", cap =", cap(slice2))

	// 1.3 创建切片方式3
	slice3 := [...]int{11, 12, 13}
	fmt.Println("slice3 =", slice3, ", len =", len(slice3), ", cap =", cap(slice3))
}

func query(nums []int) {
	// 2.0 通过切片访问数组元素
	fmt.Println("\nall =", nums[:])
	fmt.Println("\ttop 3 =", nums[:3])
	fmt.Println("\tafter top 3 =", nums[3:])
	fmt.Println("\tafter 3 =", nums[len(nums)-3:])
}
