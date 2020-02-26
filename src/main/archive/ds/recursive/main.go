package main

import "fmt"

// 简单递归 知其真意
func main() {
	n := 4
	test1(n)

	test2(n)
}

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println(n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println(n)
	}
}
