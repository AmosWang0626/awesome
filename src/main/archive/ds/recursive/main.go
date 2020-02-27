package main

import "fmt"

/*
递归需要遵循的重要原则:
1.执行一个函数时，就创建一个新的受保护的独立空间（新函数栈）
2.函数的局部变量是独立的，不会相互影响。如果希望各个函数栈使用
  同一个数据，一般使用引用传递。
3.递归必须向退出递归的条件逼近，否则就是无限递归，死龟了：)
4.当一个函数执行完毕，或者遇到return就会返回，遵守谁调用，就将
  结果返回给谁，同时当函数执行完毕或者return时，该函数本身也会
  被系统销毁。

递归可以解决什么样的问题
1.各种数学问题：8皇后问题，汉诺塔，阶乘问题，迷宫问题，球和篮子问题(Google编程大赛)
2.将用栈解决的问题 >>> 递归代码比较简洁
*/

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
