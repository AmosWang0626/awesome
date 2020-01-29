package main

import "fmt"

var age = test()

func test() int {
	fmt.Println("init age")
	return 18
}

func init() {
	fmt.Println("base init")
}

func main() {
	fmt.Println("base main", age)
}
