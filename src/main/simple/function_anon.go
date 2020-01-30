package main

import (
	"fmt"
	"strings"
)

var baseFun = func(n1, n2 int) int {
	return n1 + n2
}

func main() {
	testAnon()
	fmt.Println(".......................")
	ada := Adapter()
	for i := 0; i < 3; i++ {
		fmt.Println(ada(64))
	}
	fmt.Println(".......................")

	fmt.Println(strings.HasPrefix("dev-hello", "dev"))
	fmt.Println(strings.HasSuffix("dev-hello.go", ".go"))
}

func Adapter() func(int) int {
	n := 0
	str := "hello, "
	return func(i int) int {
		n += i
		str += string(i)
		fmt.Println(str)
		return n
	}
}

// 匿名函数测试
func testAnon() {
	var f1 = func(n1, n2 int) int {
		return n1 + n2
	}
	fmt.Println("f1", f1(100, 200))

	var f2 = func(n1, n2 int) int {
		return n1 + n2
	}(111, 22)
	fmt.Println("f2", f2)

	// 全局匿名函数
	fmt.Println("baseFun", baseFun(50, 60))
}
