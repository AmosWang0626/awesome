package main

import "fmt"

func main() {
	n1 := 100
	n2 := 125
	// %T 打印类型
	fmt.Printf("n1 type=%T\n", n1)

	fmt.Println("add1(n1, n2) >>>>>>>>>>", add1(n1, n2))

	fmt.Println(add2(n1, n2))
	_, _, sum := add2(n1, n2)
	fmt.Println("_, _, sum := add2(n1, n2) >>>>>>>>>>", sum)

	fmt.Println(add3(n1, n2))

	fmt.Println("second(add1, 1, 2) >>>>>>>>>>", second(add1, 1, 2))
	fmt.Println("second2(add1, 1, 2) >>>>>>>>>>", second2(add1, 11, 22))

	fmt.Println("Recursive Call >>>>>>>>>>", reCall(0))
}

func add1(n1 int, n2 int) int {
	return n1 + n2
}

func add2(n1 int, n2 int) (int, int, int) {
	return n1, n2, n1 + n2
}
func add3(n1 int, n2 int) (msg string, sum int) {
	msg = fmt.Sprintf("[input %d, %d] >>> ", n1, n2)
	sum = n1 + n2
	return
}

func second(fun func(n1 int, n2 int) int, n1 int, n2 int) int {
	return fun(n1, n2)
}

type funType = func(n1 int, n2 int) int

func second2(fun funType, n1 int, n2 int) int {
	return fun(n1, n2)
}

func reCall(num int) int {
	num++
	if num < 10 {
		return reCall(num)
	}
	return num
}
