package main

import "fmt"

func main() {
	fmt.Println("444 sum =", sum(100, 200))
	fmt.Println(".........end.............")

	// monitor db run sql
	monitor()
}

func sum(n1, n2 int) int {
	// defer 会将操作依次入栈 先进后出
	defer fmt.Println("111 n1 =", n1)
	defer fmt.Println("111 n2 =", n2)

	res := n1 + n2
	fmt.Println("222 res =", res)
	n1++
	n2++
	res = n1 + n2
	fmt.Printf("333 n1=%d, n2=%d, res=%d \n", n1, n2, res)

	return res
}

// defer 最佳实践
func monitor() {
	fmt.Println("1.1 加载驱动,建立数据库连接")
	defer fmt.Println("1.2 关闭数据库连接")

	fmt.Println("\t2.1 准备运行环境")
	defer fmt.Println("\t2.2 关闭运行环境")

	fmt.Println("\t\t3.1 SQL线程开始执行")
	defer fmt.Println("\t\t3.2 关闭SQL线程")

	fmt.Println("\t\t\t4.0 处理运行结果")
}
