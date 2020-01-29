package main

import "fmt"

func main() {
	fmt.Println("555555555555 sum=", sum(100, 200))
	fmt.Println("666666666666")
}

func sum(n1, n2 int) int {
	defer fmt.Println("111111111111")
	defer fmt.Println("222222222222")
	defer fmt.Println("333333333333")

	fmt.Println("444444444444")

	return n1 + n2
}
